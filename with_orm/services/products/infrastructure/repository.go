package infrastructure

import (
	"gorm.io/gorm"
	errorUtils "with.orm/libs/error-utils"
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
	err := repository.db.Save(product).Error
	return errorUtils.Wrap(err)
}

func (repository *productRepositoryImpl) FindOne(id string) (*domain.Product, error) {
	product := &domain.Product{}
	err := repository.db.Where("id = ?", id).First(product).Error
	return product, errorUtils.Wrap(err)
}

func (repository *productRepositoryImpl) Find() ([]*domain.Product, error) {
	products := []*domain.Product{}
	err := repository.db.Find(&products).Error
	return products, errorUtils.Wrap(err)
}
