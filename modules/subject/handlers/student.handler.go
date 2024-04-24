package subject_handlers

import (
	subject_models "lms_backend/modules/subject/models"
	"lms_backend/session"

	"github.com/craftzbay/go_grc/v2/helpers"
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
