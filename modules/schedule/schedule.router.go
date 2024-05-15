package schedule

import (
	schedule_handlers "lms_backend/modules/schedule/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var scheduleHandler schedule_handlers.ScheduleHandler

	scheduleApi := app.Group("schedule")
	scheduleApi.Get("", scheduleHandler.List)
}
