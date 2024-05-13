package subject

import (
	subject_handlers "lms_backend/modules/subject/handlers"
	"lms_backend/session"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var subjectHandler subject_handlers.SubjectHandler

	subjectApi := app.Group("subject")
	subjectApi.Get("", session.TokenMiddleware, subjectHandler.List)
	subjectApi.Post("", session.TokenMiddleware, subjectHandler.Create)
	subjectApi.Put("", session.TokenMiddleware, subjectHandler.Update)
	subjectApi.Delete("", session.TokenMiddleware, subjectHandler.Delete)

	var taskHandler subject_handlers.TaskHandler

	taskApi := app.Group("task", session.TokenMiddleware)
	taskApi.Get("", taskHandler.List)
	taskApi.Post("", taskHandler.Create)
	taskApi.Put("", taskHandler.Update)
	taskApi.Delete("", taskHandler.Delete)

	var examHandler subject_handlers.ExamHandler

	examApi := app.Group("exam", session.TokenMiddleware)
	examApi.Get("", examHandler.List)
	examApi.Post("", examHandler.Create)
	examApi.Put("", examHandler.Update)
	examApi.Delete("", examHandler.Delete)

	var studentHandler subject_handlers.StudentHandler

	studentApi := app.Group("student", session.TokenMiddleware)
	studentApi.Get("subject", studentHandler.StudentSubject)
	studentApi.Get("task", studentHandler.StudentTask)
}
