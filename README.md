# go-drip
Drip API wrapper for golang

## Features

- Create a subscriber by email address
- Track an event
- Custom fields for subscribers and events
- Tagging/untagging subscribers
- Custom field [key normalization](#key-normailzation)

Not yet implemented:

- [ ] Other API endpoints

## Examples

```go
dripClient = drip.NewClient("DRIP_API_TOKEN", "DRIP_ACCOUNT_ID")
_ = dripClient.CreateSubscriber("an@email-adress.com", map[string]interface{}{ "first_name": "Chris" })
_ = dripClient.RecordEvent("an@email-address.com", "Registered")
_ = dripClient.TagSubscriber("an@email-address.com", "Beta")
```

## Key normailzation

`go-drip` will normalize custom fields keys by removing `$` characters,
replacing spaces with underscores, and lower casing the entire key. For
example:

```
"First Name" => "first_name"
"$email" => "email"
```

Drip will not add any custom fields even if only one key is malformed. This
feature tries to guard you against that and helps avoid re-implementing the
same transformations. I'm happy to accept or add additional transformations, if
you'd benefit from them.
