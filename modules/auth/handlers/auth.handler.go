package auth_handlers

import (
	auth_models "lms_backend/modules/auth/models"
	"lms_backend/modules/otp"
	user_models "lms_backend/modules/user/models"
	"lms_backend/session"
	"lms_backend/utils"

	"lms_backend/helpers"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func (*AuthHandler) Login(c *fiber.Ctx) error {
	req := new(auth_models.ReqLogin)

	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if err := gvalidate.Validate(*req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	user, err := user_models.FindUserBy("username", req.Username)
	if err != nil {
		return helpers.ResponseBadRequest(c, "User not found")
	}

	if user.Password != helpers.GeneratePassword(req.Password) {
		return helpers.ResponseBadRequest(c, "Wrong username or password")
	}

	tokenInfo := new(session.Token)
	tokenInfo.SetUserId(user.Id)

	token, err := session.GetToken(tokenInfo)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, session.ResToken{Token: token})
}

func (*AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	req := new(auth_models.ReqForgot)

	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*req); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}
	if err := otp.CheckOtp(req.Email, req.Otp); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	user, err := user_models.FindUserBy("email", req.Email)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	user.Password = helpers.GeneratePassword(req.Password)

	if err := user.Update(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}

func (*AuthHandler) GetOtp(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return helpers.ResponseBadRequest(c, "email field is required")
	}

	if !gvalidate.IsEmail(email) {
		return helpers.ResponseBadRequest(c, "Invalid email")
	}

	if !utils.CheckExists("user", []string{"email"}, []interface{}{email}) {
		return helpers.ResponseBadRequest(c, "Email address not found")
	}

	if err := otp.SendOtp(email); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	return helpers.Response(c, "OTP code sent")
}
