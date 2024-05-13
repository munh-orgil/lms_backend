package subject_handlers

import (
	"lms_backend/helpers"
	subject_models "lms_backend/modules/subject/models"
	"lms_backend/session"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct{}

func (*StudentHandler) StudentSubject(c *fiber.Ctx) error {
	subjectId := uint(c.QueryInt("subject_id"))
	studentId := session.GetTokenInfo(c).GetUserId()
	res, err := subject_models.GetSubject(studentId, subjectId)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, res)
}

func (*StudentHandler) StudentTask(c *fiber.Ctx) error {
	studentId := session.GetTokenInfo(c).GetUserId()
	res, err := subject_models.GetTasks(studentId)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, res)
}
