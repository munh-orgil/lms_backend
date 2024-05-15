package schedule_handlers

import (
	"lms_backend/helpers"
	schedule_models "lms_backend/modules/schedule/models"

	"github.com/gofiber/fiber/v2"
)

type ScheduleHandler struct{}

func (*ScheduleHandler) List(c *fiber.Ctx) error {
	res, err := schedule_models.ScheduleList()
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, res)
}
