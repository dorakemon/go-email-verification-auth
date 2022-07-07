package controllers

import (
	"api-service/cmd/server/helpers"
	"api-service/cmd/server/models"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func VerifyEmailController(c echo.Context) error {
	email := c.FormValue("email")
	err := helpers.VerifyEmailScheme(email)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<strong>Hello, World!</strong>")
	}
	otp, err := helpers.GenerateOtp()
	fmt.Printf("otp: %s", otp)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "<strong>cannot create otp</strong>")
	}

	sessionKey := helpers.GenerateSessionKey()

	// rabbitmqに登録を行う

	// err = sendOtpPasswordMail(otp)
	// if err != nil {
	// 	return c.HTML(http.StatusInternalServerError, "<strong>cannot send email</strong>")
	// }

	models.SetSession(sessionKey, email, otp)

	cookie := &http.Cookie{
		Name:    "email_verification",
		Value:   sessionKey,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
		// Domain:  "example.com",
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/check-otp")
}
