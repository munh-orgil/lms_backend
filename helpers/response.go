package helpers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Response(c *fiber.Ctx, responseDetails ...interface{}) error {

	if len(responseDetails) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success"})
	}

	var res interface{}
	switch val := responseDetails[0].(type) {
	case string:
		res = fiber.Map{"message": val}
	default:
		res = val
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func ResponseForbidden(c *fiber.Ctx) error {
	msg := "Unauthorized"
	zap.L().Error(fiber.ErrForbidden.Message, zap.String("error", msg))
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": msg})
}

func ResponseUnauthorized(c *fiber.Ctx) error {
	msg := "Invalid token"
	zap.L().Error(fiber.ErrUnauthorized.Message, zap.String("error", msg))
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": msg})
}

func ResponseBadRequest(c *fiber.Ctx, message string) error {
	zap.L().Error(fiber.ErrBadRequest.Message, zap.String("error", message))
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

func ResponseErr(c *fiber.Ctx, message string) error {
	zap.L().Error(fiber.ErrInternalServerError.Message, zap.String("error", message))
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}
