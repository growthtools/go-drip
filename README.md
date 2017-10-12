# go-drip
Drip API wrapper for golang

## Features

- Create a subscriber by email address
- Track an event

Not yet implemented:

- [ ] Custom fields for subscribers and events
- [ ] Other API endpoints

## Examples

```go
dripClient = drip.NewClient("DRIP_API_TOKEN", "DRIP_ACCOUNT_ID")
_ = dripClient.CreateSubscriber("an@email-adress.com")
_ = dripClient.RecordEvent("an@email-address.com", "Registered")
```
