package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// create a fiber application
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start the application
	app.Listen(":3000")
}
