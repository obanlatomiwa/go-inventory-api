package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/obanlatomiwa/go-inventory-api/models"
	"time"
)

var storage []models.Item

func getAllItems() []models.Item {
	return storage
}

func getItemById(id string) (models.Item, error) {
	for _, item := range storage {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item not found")
}

func createItem(itemRequest models.ItemRequest) models.Item {
	newItem := models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	storage = append(storage, newItem)

	return newItem
}

func updateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {

	for index, item := range storage {
		if item.ID == id {
			item.Name = itemRequest.Name
			item.Price = itemRequest.Price
			item.Quantity = itemRequest.Quantity
			item.UpdatedAt = time.Now()

			storage[index] = item
			return item, nil
		}
	}
	return models.Item{}, errors.New("Item update failed, Item not found")
}

func deleteItemById(id string) bool {
	for index, item := range storage {
		if item.ID == id {
			storage = append(storage[:index], storage[index+1:]...)
			return true
		}
	}
	return false
}
