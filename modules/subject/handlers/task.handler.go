package subject_handlers

import (
	"lms_backend/helpers"
	subject_models "lms_backend/modules/subject/models"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct{}

func (*TaskHandler) List(c *fiber.Ctx) error {
	if res, err := subject_models.TaskList(c); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*TaskHandler) Create(c *fiber.Ctx) error {
	data := new(subject_models.Task)
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

func (*TaskHandler) Update(c *fiber.Ctx) error {
	data := new(subject_models.Task)
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

func (*TaskHandler) Delete(c *fiber.Ctx) error {
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
	subject := new(subject_models.Task)
	subject.Id = req.Id
	if err := subject.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
