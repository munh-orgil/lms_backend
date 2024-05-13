package attachment

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	var attachmentHandler AttachmentHandler

	attachmentApi := app.Group("attachment")
	attachmentApi.Get("", attachmentHandler.List)
	attachmentApi.Post("", attachmentHandler.Create)
	attachmentApi.Put("", attachmentHandler.Update)
	attachmentApi.Delete("", attachmentHandler.Delete)
}
