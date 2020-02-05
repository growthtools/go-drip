package drip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const baseURL = "https://api.getdrip.com/v2/"

// Client represents a Drip API client wrapper
type Client struct {
	apiKey     string
	appID      string
	httpClient http.Client
}

type subParams struct {
	Email        string                 `json:"email"`
	CustomFields map[string]interface{} `json:"custom_fields"`
}

type subRoot struct {
	Subscribers []subParams `json:"subscribers"`
}

type eventRoot struct {
	Events []eventParams `json:"events"`
}

type eventParams struct {
	Email  string `json:"email"`
	Action string `json:"action"`
}

type tagRoot struct {
	Tags []tagParams `json:"tags"`
}

type tagParams struct {
	Email string `json:"email"`
	Tag   string `json:"tag"`
}

// NewClient returns a client instance ready to act with Drip for the given app and API key
func NewClient(apiKey, appID string) *Client {
	client := &Client{
		apiKey: apiKey,
		appID:  appID,
		httpClient: http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: time.Second * 5,
				}).Dial,
				TLSHandshakeTimeout: time.Second * 5,
			},
		},
	}

	return client
}

// CreateSubscriber creates a new or updates an existing subscriber by email
func (c Client) CreateSubscriber(email string, customFields map[string]interface{}) error {
	dripFields := map[string]interface{}{}
	for key, value := range customFields {
		newKey := key
		newKey = strings.Replace(key, "$", "", -1)
		newKey = strings.Replace(newKey, " ", "_", -1)
		newKey = strings.ToLower(newKey)
		dripFields[newKey] = value
	}
	bodyData := subRoot{
		Subscribers: []subParams{
			{Email: email, CustomFields: dripFields},
		},
	}

	return c.authenticatedPost("/subscribers", bodyData)
}

// RecordEvent sends a custom event to Drip
func (c Client) RecordEvent(email, eventName string) error {
	bodyData := eventRoot{
		Events: []eventParams{
			{Email: email, Action: eventName},
		},
	}
	return c.authenticatedPost("/events", bodyData)
}

// TagSubscriber adds a tag to a subscriber
func (c Client) TagSubscriber(email, tag string) error {
	data := tagRoot{
		Tags: []tagParams{
			{Email: email, Tag: tag},
		},
	}
	return c.authenticatedPost("/tags", data)
}

func (c Client) authenticatedPost(path string, body interface{}) error {
	postBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := c.authenticatedRequest("POST", path, bytes.NewReader(postBody))
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, err := ioutil.ReadAll(resp.Body)
		defer func() { _ = resp.Body.Close() }()
		if err != nil {
			return err
		}
		return fmt.Errorf("Drip API error: (%d} %s)", resp.StatusCode, string(body))
	}

	return nil
}

func (c Client) authenticatedRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", baseURL+c.appID+path, body)
	req.SetBasicAuth(c.apiKey, "")
	req.Header.Add("Accept", "application/vnd.api+json")

	return req, err
}
