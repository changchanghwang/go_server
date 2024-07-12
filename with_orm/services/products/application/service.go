package application

import (
	product "with.orm/services/products/domain"
	"with.orm/services/products/infrastructure"
)

type ProductService struct {
	productRepository infrastructure.ProductRepository
}

func NewProductService(productRepository infrastructure.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}

func (service *ProductService) Create(name string) error {
	product := product.New(name)
	return service.productRepository.Save(product)
}
func (service *ProductService) List() ([]*product.Product, error) {
	return service.productRepository.Find()
}
