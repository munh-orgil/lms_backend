package subject_handlers

import (
	"lms_backend/helpers"
	subject_models "lms_backend/modules/subject/models"
	"lms_backend/session"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct{}

func (*SubjectHandler) List(c *fiber.Ctx) error {
	studentId := session.GetTokenInfo(c).GetUserId()
	res, err := subject_models.SubjectList(studentId)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, res)
}

func (*SubjectHandler) Create(c *fiber.Ctx) error {
	data := new(subject_models.Subject)
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

func (*SubjectHandler) Update(c *fiber.Ctx) error {
	data := new(subject_models.Subject)
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

func (*SubjectHandler) Delete(c *fiber.Ctx) error {
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
	subject := new(subject_models.Subject)
	subject.Id = req.Id
	if err := subject.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
