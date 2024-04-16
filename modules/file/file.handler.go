package file

import (
	"path/filepath"

	"lms_backend/helpers"

	"github.com/gofiber/fiber/v2"
)

type FileHandler struct{}

const (
	DISALLOWED_FILE_SIZE = "file exceeded the allowed size"
	MAX_FILE_SIZE        = 5000 * 1024 * 1024
	ROOT_DIR             = "./data"
)

func (*FileHandler) Find(c *fiber.Ctx) error {
	file := c.Query("file")
	filePath := filepath.Join(ROOT_DIR, file)
	return c.SendFile(filePath, false)
}

func (*FileHandler) Upload(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	name, err := UploadFile(fileHeader, c)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c, fiber.Map{
		"name": name,
	})
}
