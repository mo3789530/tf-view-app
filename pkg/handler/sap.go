package handler

import "github.com/gofiber/fiber/v2"

func Sap(c *fiber.Ctx) error {
	return c.SendFile("ui/build/index.html")
}
