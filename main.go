package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/database"
	"github.com/obanlatomiwa/go-inventory-api/routes"
	"github.com/obanlatomiwa/go-inventory-api/utils"
)

const DefaultPort = "3000"

func NewApp() *fiber.App {
	// create a fiber application
	app := fiber.New()

	routes.SetUpRoutes(app)

	// start the application
	return app
}

func main() {
	PORT := utils.GetValueFromConfigFile("APP_PORT")

	if PORT == "" {
		PORT = DefaultPort
	}
	app := NewApp()
	database.InitialiseDatabase(utils.GetValueFromConfigFile("DB_NAME"))

	app.Listen(fmt.Sprintf(":%s", PORT))

}
