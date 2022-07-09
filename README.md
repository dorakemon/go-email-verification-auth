# Golang Authentication Email Verification

## description

user authentication service written by go

using...
- email&password auth
- verify email address by sending One Time Password (OTP)
- RabbitMQ for message broker
- 2 microservices

## setup

### add `email-services/.env`

```sh
FROM_MAIL=[gmail address for sending something]
GMAIL_PASS=[gmail password]
```

### build

```sh
make up_build
```

## down

```sh
make down
```

## Todo

- [x] setup rabbitmq
- [x] email service
