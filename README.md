# pingdom-twilio-integration

[![Publish Docker](https://github.com/DocPlanner/pingdom-twilio-integration/workflows/Publish%20Docker/badge.svg)](https://hub.docker.com/r/docplanner/pingdom-twilio-integration)

## How it works?
pingdom-twilio-integration sends messages to contact groups defined in configuration.

## Pingdom configuration
Only thing to do is to add webhook to your check in Pingdom

### Sample webhook url
`https://your.domain.tld/?secret=SECRET&contact_group=ops`
SECRET is a value defined in config file as secret

## How to use
To use this tool run `pingdom-twilio-integration` with `-config config.yaml`
```yaml
contacts:
  bob: "+48999888777"
  tod: "+48888777666"
contact_groups:
  ops: [bob, tod]
twilio:
  phone_from: "+48444555666"
  account_sid: "xxx"
  auth_token: "xxx"
secret: "SECRET"
```
