package infrastructure

import (
	"gorm.io/gorm"
	errorUtils "with.orm/libs/error-utils"
	"with.orm/services/inventories/domain"
)

type InventoryRepository interface {
	Save(inventory *domain.Inventory) error
	FindOne(id string) (*domain.Inventory, error)
	Find() ([]*domain.Inventory, error)
}

type inventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepositoryImpl{db}
}

func (repository *inventoryRepositoryImpl) Save(inventory *domain.Inventory) error {
	err := repository.db.Save(inventory).Error
	return errorUtils.Wrap(err)
}

func (repository *inventoryRepositoryImpl) FindOne(id string) (*domain.Inventory, error) {
	inventory := &domain.Inventory{}
	err := repository.db.Where("id = ?", id).First(inventory).Error
	return inventory, errorUtils.Wrap(err)
}

func (repository *inventoryRepositoryImpl) Find() ([]*domain.Inventory, error) {
	inventorys := []*domain.Inventory{}
	err := repository.db.Find(&inventorys).Error
	return inventorys, errorUtils.Wrap(err)
}
