package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/handlers"
)

func SetUpRoutes(app *fiber.App) {
	// register auth handlers
	app.Post("/api/v1/signup", handlers.SignUp)
	app.Post("/api/v1/login", handlers.Login)

	app.Get("/api/v1/items", handlers.GetAllItems)
	app.Get("/api/v1/items/:id", handlers.GetItemById)
	app.Post("/api/v1/items", handlers.CreateItem)
	app.Put("/api/v1/items/:id", handlers.UpdateItem)
	app.Delete("/api/v1/items", handlers.DeleteItem)
}
