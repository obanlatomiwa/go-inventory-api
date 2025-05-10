package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/routes"
)

func main() {
	// create a fiber application
	app := fiber.New()

	routes.SetUpRoutes(app)

	// start the application
	app.Listen(":3000")
}
