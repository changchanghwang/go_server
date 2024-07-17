package application

import (
	inventory "with.orm/services/inventories/domain"
	"with.orm/services/inventories/infrastructure"
)

type InventoryService struct {
	inventoryRepository infrastructure.InventoryRepository
}

func NewInventoryService(inventoryRepository infrastructure.InventoryRepository) *InventoryService {
	return &InventoryService{inventoryRepository}
}

func (service *InventoryService) Create(productId string, stock int) error {
	inventory := inventory.New(productId, stock)
	return service.inventoryRepository.Save(inventory)
}

func (service *InventoryService) List() ([]*inventory.Inventory, error) {
	return service.inventoryRepository.Find()
}
