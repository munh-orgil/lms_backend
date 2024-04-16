package auth

import (
	auth_handlers "lms_backend/modules/auth/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	authHandler := auth_handlers.AuthHandler{}
	authApi := app.Group("auth")
	authApi.Get("otp", authHandler.GetOtp)
	authApi.Post("login", authHandler.Login)
	authApi.Post("forgot", authHandler.ForgotPassword)
}
