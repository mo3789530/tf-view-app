package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("state")
	if err != nil {
		return err
	}

	return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
}
