# Golang Authentication Email Verification

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
