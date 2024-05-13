package subject_handlers

import (
	"lms_backend/helpers"
	subject_models "lms_backend/modules/subject/models"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type ExamHandler struct{}

func (*ExamHandler) List(c *fiber.Ctx) error {
	subjectId := uint(c.QueryInt("subject_id"))
	if res, err := subject_models.ExamList(subjectId); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*ExamHandler) Create(c *fiber.Ctx) error {
	data := new(subject_models.Exam)
	if err := c.BodyParser(data); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*data); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	if err := data.Create(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*ExamHandler) Update(c *fiber.Ctx) error {
	data := new(subject_models.Exam)
	if err := c.BodyParser(data); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	if errors := gvalidate.Validate(*data); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	if err := data.Update(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*ExamHandler) Delete(c *fiber.Ctx) error {
	type Req struct {
		Id uint `json:"id" validate:"required"`
	}
	req := new(Req)
	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if err := gvalidate.Validate(*req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	subject := new(subject_models.Exam)
	subject.Id = req.Id
	if err := subject.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
