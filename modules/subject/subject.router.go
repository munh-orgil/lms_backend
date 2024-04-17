package subject

import (
	subject_handlers "lms_backend/modules/subject/handlers"
	"lms_backend/session"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var subjectHandler subject_handlers.SubjectHandler

	subjectApi := app.Group("subject", session.TokenMiddleware)
	subjectApi.Get("", subjectHandler.List)
	subjectApi.Post("", subjectHandler.Create)
	subjectApi.Put("", subjectHandler.Update)
	subjectApi.Delete("", subjectHandler.Delete)

	var taskHandler subject_handlers.TaskHandler

	taskApi := app.Group("task")
	taskApi.Get("", taskHandler.List)
	taskApi.Post("", taskHandler.Create)
	taskApi.Put("", taskHandler.Update)
	taskApi.Delete("", taskHandler.Delete)
}
