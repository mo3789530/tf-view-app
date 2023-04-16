package main

import (
	"log"
	"tfview/pkg/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Static("", "ui/build")
	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
