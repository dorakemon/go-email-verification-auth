package models

import "errors"

type SessionValue struct {
	Email string
	Otp   string
}

var (
	SessionOtpMap = map[string]SessionValue{}
)

func AddSession(session string, email string, otp string) {
	SessionOtpMap[session] = SessionValue{
		Email: email,
		Otp:   otp,
	}
}

func GetSession(session string) (SessionValue, error) {
	sessionValue, found := SessionOtpMap[session]
	if !found {
		return sessionValue, errors.New("session was expired")
	}
	return sessionValue, nil
}

func SetSession(session string, email string, otp string) {
	SessionOtpMap[session] = SessionValue{
		Email: email,
		Otp:   otp,
	}
}
