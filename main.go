package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("", "ui/build")
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("ui/build/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}
