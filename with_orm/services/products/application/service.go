package application

import (
	errorUtils "with.orm/libs/error-utils"
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
	err := service.productRepository.Save(product)
	return errorUtils.WrapWithCode(err, 500)
}
func (service *ProductService) List() ([]*product.Product, error) {
	products, err := service.productRepository.Find()
	return products, errorUtils.WrapWithCode(err, 500)
}
