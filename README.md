# go-drip
Drip API wrapper for golang

## Features

- Create a subscriber by email address
- Track an event
- Custom fields for subscribers and events
- Tagging subscribers

Not yet implemented:

- [ ] Other API endpoints

## Examples

```go
dripClient = drip.NewClient("DRIP_API_TOKEN", "DRIP_ACCOUNT_ID")
_ = dripClient.CreateSubscriber("an@email-adress.com", map[string]interface{}{ "Partnerships": "Chris" })
_ = dripClient.RecordEvent("an@email-address.com", "Registered")
_ = dripClient.TagSubscriber("an@email-address.com", "Beta")
```
