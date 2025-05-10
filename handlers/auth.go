package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/services"

	"net/http"
)

func SignUp(c *fiber.Ctx) error {
	var userInput *models.UserRequest
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := userInput.ValidateUserRequest()
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}

	jwt, err := services.Signup(*userInput)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(models.Response[any]{
		Success: true,
		Message: "token generated",
		Data:    jwt,
	})
}

func Login(c *fiber.Ctx) error {
	var userInput *models.UserRequest
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := userInput.ValidateUserRequest()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}

	jwt, err := services.Login(*userInput)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "token generated",
		Data:    jwt,
	})
}
