package infrastructure

import (
	"gorm.io/gorm"
	"with.orm/services/products/domain"
)

type ProductRepository interface {
	Save(product *domain.Product) error
	FindOne(id string) (*domain.Product, error)
	Find() ([]*domain.Product, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db}
}

func (repository *productRepositoryImpl) Save(product *domain.Product) error {
	return repository.db.Save(product).Error
}

func (repository *productRepositoryImpl) FindOne(id string) (*domain.Product, error) {
	product := &domain.Product{}
	err := repository.db.Where("id = ?", id).First(product).Error
	return product, err
}

func (repository *productRepositoryImpl) Find() ([]*domain.Product, error) {
	products := []*domain.Product{}
	err := repository.db.Find(&products).Error
	return products, err
}
