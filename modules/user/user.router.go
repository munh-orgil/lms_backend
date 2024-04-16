package user

import (
	user_handlers "lms_backend/modules/user/handlers"
	"lms_backend/session"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	userHandler := user_handlers.UserHandler{}
	userApi := app.Group("user")
	userApi.Get("", session.TokenMiddleware, userHandler.List)
	userApi.Get("profile", session.TokenMiddleware, userHandler.Profile)
	userApi.Post("", session.TokenMiddleware, userHandler.Create)
	userApi.Put("", session.TokenMiddleware, userHandler.Update)
}
