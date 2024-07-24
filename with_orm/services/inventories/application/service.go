package application

import (
	errorUtils "with.orm/libs/error-utils"
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
	err := service.inventoryRepository.Save(inventory)
	return errorUtils.WrapWithCode(err, 500)
}

func (service *InventoryService) List() ([]*inventory.Inventory, error) {
	inventories, err := service.inventoryRepository.Find()
	return inventories, errorUtils.WrapWithCode(err, 500)
}
