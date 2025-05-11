package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/handlers"
	"github.com/obanlatomiwa/go-inventory-api/middlewares"
)

func SetUpRoutes(app *fiber.App) {
	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")
	publicRoutes.Post("/signup", handlers.SignUp)
	publicRoutes.Post("/login", handlers.Login)

	// private routes, authentication required
	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())

	privateRoutes.Get("/items", handlers.GetAllItems)
	privateRoutes.Get("/items/:id", handlers.GetItemById)
	privateRoutes.Post("/items", handlers.CreateItem)
	privateRoutes.Put("/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", handlers.DeleteItem)
}
