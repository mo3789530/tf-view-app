package router

import (
	"tfview/pkg/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
	})

	api.Post("/tf/state", )

	app.Get("/*", handler.Sap)
}
