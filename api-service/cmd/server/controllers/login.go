package controllers

import (
	"api-service/cmd/server/models"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

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

	claims := &jwtCustomClaims{
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// must be scret
	// this private key is sample
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<strong>password or email is wrong</strong>")
	}

	content := fmt.Sprintf("<div><h2>Logged in sccessfully</h2><h3>Your JWT Token is %s</h3></div>", t)

	return c.HTML(http.StatusBadRequest, content)
}
