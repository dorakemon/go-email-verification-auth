package controllers

import (
	"api-service/cmd/server/helpers"
	"api-service/cmd/server/models"
	"api-service/event"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (config *Config) VerifyEmailController(c echo.Context) error {
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
	producer, err := event.NewEventProducer(config.Conn)
	if err != nil {
		log.Println(err)
		return c.HTML(http.StatusInternalServerError, "<strong>cannot create otp</strong>")
	}
	content := fmt.Sprintf("Verification Code is \n%s", otp)
	payload := EmailPayload{
		To:      email,
		Subject: "Verification Code",
		Content: content,
	}
	j, err := json.MarshalIndent(&payload, "", "\t")
	if err != nil {
		log.Println(err)
		return c.HTML(http.StatusInternalServerError, "<strong>cannot create otp</strong>")
	}
	err = producer.Push(string(j), "email.register")
	if err != nil {
		log.Println(err)
		return c.HTML(http.StatusInternalServerError, "<strong>cannot create otp</strong>")

	}

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
