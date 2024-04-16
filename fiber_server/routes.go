package fiber_server

import (
	"lms_backend/modules/auth"
	"lms_backend/modules/file"
	"lms_backend/modules/subject"
	"lms_backend/modules/user"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	auth.SetRoutes(app)
	file.SetRoutes(app)
	user.SetRoutes(app)
	subject.SetRoutes(app)
}
