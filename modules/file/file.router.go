package file

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	fileHandler := FileHandler{}

	fileApi := app.Group("file")
	fileApi.Get("", fileHandler.Find)
	fileApi.Post("", fileHandler.Upload)
}
