package router

import (
	"api-service/cmd/server/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) *echo.Echo {

	e.GET("/email-verifier", func(c echo.Context) error {
		return c.Render(http.StatusOK, "ver-email.go.html", map[string]interface{}{
			"message": "Please write correct email",
		})
	})

	e.POST("/email-verifier", controllers.VerifyEmailController)

	e.GET("/check-otp", func(c echo.Context) error {
		return c.Render(http.StatusOK, "input-otp.go.html", map[string]interface{}{
			"message": "Please write otp from email",
		})
	})

	e.POST("/check-otp", controllers.CheckOtpController)

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register-user.go.html", map[string]interface{}{
			"message": "Register User",
		})
	})
	e.POST("/register", controllers.RegisterController)

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.go.html", map[string]interface{}{
			"message": "Login User",
		})
	})
	e.POST("/login", controllers.LoginController)

	return e
}
