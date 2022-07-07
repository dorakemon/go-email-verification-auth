package controllers

import (
	"api-service/cmd/server/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	cookie, err := c.Cookie("email_verification")
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusBadRequest, "<strong>session was expired</strong>")
	}
	sessionValue, err := models.GetSession(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusBadRequest, "<strong>session was expired</strong>")
	}
	if sessionValue.Email != email {
		return c.HTML(http.StatusBadRequest, "<strong>different from verified email</strong>")
	}

	models.AddUser(username, password, email)

	return c.HTML(http.StatusOK, "<strong>registerd</strong>")
}
