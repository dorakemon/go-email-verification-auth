package controllers

import (
	"api-service/cmd/server/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (config *Config) LoginController(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.FindUserByEmail(email)

	if err != nil {
		return c.HTML(http.StatusBadRequest, "<strong>password or email is wrong</strong>")
	}

	if user.Password != password {
		return c.HTML(http.StatusBadRequest, "<strong>password or email is wrong</strong>")
	}

	return c.HTML(http.StatusBadRequest, "<strong>logged in successfully</strong>")
}
