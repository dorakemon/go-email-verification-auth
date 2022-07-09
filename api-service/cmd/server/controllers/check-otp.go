package controllers

import (
	"api-service/cmd/server/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (config *Config) CheckOtpController(c echo.Context) error {
	otp := c.FormValue("otp")
	cookie, err := c.Cookie("email_verification")
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusBadRequest, "<strong>session was expired</strong>")
	}
	sessionValue, err := models.GetSession(cookie.Value)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<strong>session was expired</strong>")
	}
	if sessionValue.Otp != otp {
		return c.HTML(http.StatusBadRequest, "<strong>otp is wrong</strong>")
	}

	return c.Redirect(http.StatusFound, "/register")
}
