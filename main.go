package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	apiV1 := app.Group("v1")

	apiV1.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "oke boss",
		})
	})

	apiV1.Post("/test", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "oke boss",
		})
	})

	app.Listen(":3000")
}
