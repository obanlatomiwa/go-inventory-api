package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/obanlatomiwa/go-inventory-api/database"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"time"
)

var storage []models.Item

func GetAllItems() []models.Item {
	var items []models.Item
	database.DB.Order("created_at desc").Find(&items)
	return items
}

func GetItemById(id string) (models.Item, error) {
	var item models.Item
	database.DB.First(&item, "id = ?", id)
	return item, nil
}

func CreateItem(itemRequest models.ItemRequest) models.Item {
	newItem := models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	database.DB.Create(&newItem)

	return newItem
}

func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {

	updateItem := models.Item{
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		UpdatedAt: time.Now(),
	}

	database.DB.FirstOrCreate(&updateItem, "id = ?", id)
	return models.Item{}, errors.New("Item update failed, Item not found")
}

func DeleteItemById(id string) bool {
	database.DB.Delete(&models.Item{}, "id = ?", id)
	return true
}
