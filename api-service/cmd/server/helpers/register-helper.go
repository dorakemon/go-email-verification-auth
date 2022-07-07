package helpers

import (
	"api-service/cmd/server/configs"
	"errors"
	"strconv"

	"github.com/rs/xid"
)

func GenerateOtp() (string, error) {
	randomNum, err := generateFixedLengthRandomNum(configs.OtpLength)
	if err != nil {
		return "", errors.New("cannot generate otp")
	}
	return strconv.Itoa(randomNum), nil
}

func GenerateSessionKey() string {
	id := xid.New()
	return id.String()
}
