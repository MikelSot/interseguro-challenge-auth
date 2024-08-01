package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newFiber(errorHandler fiber.ErrorHandler) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      getApplicationName(),
		ErrorHandler: errorHandler,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: getAllowOrigins(),
		AllowMethods: getAllowMethods(),
	}))

	return app
}
