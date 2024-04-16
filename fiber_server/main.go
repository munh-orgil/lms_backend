package fiber_server

import (
	"fmt"
	"lms_backend/global"
	"lms_backend/modules/log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiber() {
	port := global.Conf.Port

	app := fiber.New()
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Request().Header.Add("X-Request-Start-Time", fmt.Sprintf("%d", time.Now().UnixMicro()))
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Done: log.FiberLogSaver,
	}))

	InitRoutes(app)
	panic(app.Listen(":" + port))
}
