package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"github.com/obanlatomiwa/go-inventory-api/services"
	"github.com/obanlatomiwa/go-inventory-api/utils"
	"net/http"
)

func GetAllItems(c *fiber.Ctx) error {
	// check the token
	err := validateToken(c)
	if err != nil {
		return err
	}

	var items []models.Item = services.GetAllItems()

	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "Successfully fetched all items",
		Data:    items,
	})
}

func GetItemById(c *fiber.Ctx) error {
	// check the token
	err := validateToken(c)
	if err != nil {
		return err
	}

	itemId := c.Params("id")
	item, err := services.GetItemById(itemId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "Successfully fetched item",
		Data:    item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	// check the token
	err := validateToken(c)
	if err != nil {
		return err
	}

	var item *models.ItemRequest
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := item.ValidateItemRequest()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}

	createdItem := services.CreateItem(*item)

	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "Successfully created item",
		Data:    createdItem,
	})
}

func UpdateItem(c *fiber.Ctx) error {
	// check the token
	err := validateToken(c)
	if err != nil {
		return err
	}

	var item *models.ItemRequest
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := item.ValidateItemRequest()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Validation error",
			Data:    errors,
		})
	}

	updatedItem, err := services.UpdateItem(*item, c.Params("id"))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "Successfully updated item",
		Data:    updatedItem,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	// check the token
	err := validateToken(c)
	if err != nil {
		return err
	}

	itemId := c.Params("id")
	deleted := services.DeleteItemById(itemId)

	if deleted {
		return c.Status(http.StatusAccepted).JSON(models.Response[any]{
			Success: true,
			Message: "Successfully deleted item",
		})
	}

	return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "Item not found",
	})
}

func validateToken(c *fiber.Ctx) error {
	// check the token
	isValid, err := utils.CheckToken(c)

	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	return nil
}
