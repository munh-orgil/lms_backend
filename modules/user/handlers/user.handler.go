package user_handlers

import (
	"lms_backend/helpers"
	user_models "lms_backend/modules/user/models"
	"lms_backend/session"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (*UserHandler) List(c *fiber.Ctx) error {
	if res, err := user_models.UserList(c); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*UserHandler) Profile(c *fiber.Ctx) error {
	userId := session.GetTokenInfo(c).GetUserId()
	user, err := user_models.FindUserBy("id", userId)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, user)
}

func (*UserHandler) Create(c *fiber.Ctx) error {
	req := new(user_models.ReqUserCreate)
	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*req); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	data := user_models.User{
		Username:       req.Username,
		Password:       req.Password,
		Email:          req.Email,
		ProfilePicture: req.ProfilePicture,
		UserType:       req.UserType,
		Lastname:       req.Lastname,
		Firstname:      req.Firstname,
		Gender:         req.Gender,
		BirthDate:      req.BirthDate,
	}

	if err := data.Create(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*UserHandler) Update(c *fiber.Ctx) error {
	req := new(user_models.ReqUserUpdate)
	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	if errors := gvalidate.Validate(*req); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	data := user_models.User{
		Id:             req.Id,
		Username:       req.Username,
		ProfilePicture: req.ProfilePicture,
		UserType:       req.UserType,
		Lastname:       req.Lastname,
		Firstname:      req.Firstname,
		Gender:         req.Gender,
		BirthDate:      req.BirthDate,
	}

	if err := data.Update(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}
