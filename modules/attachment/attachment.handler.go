package attachment

import (
	"lms_backend/helpers"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type AttachmentHandler struct{}

func (*AttachmentHandler) List(c *fiber.Ctx) error {
	if res, err := AttachmentList(c); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*AttachmentHandler) Create(c *fiber.Ctx) error {
	data := new(Attachment)
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

func (*AttachmentHandler) Update(c *fiber.Ctx) error {
	data := new(Attachment)
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

func (*AttachmentHandler) Delete(c *fiber.Ctx) error {
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
	attachment := new(Attachment)
	attachment.Id = req.Id
	if err := attachment.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
