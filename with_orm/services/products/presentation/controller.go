package presentation

import (
	"github.com/gin-gonic/gin"
	"with.orm/services/products/application"
)

type ProductController struct {
	productService *application.ProductService
}

func NewProductController(productService *application.ProductService) *ProductController {
	return &ProductController{productService}
}

func (controller *ProductController) Route(router *gin.RouterGroup) error {
	router.POST("/", controller.create)
	router.GET("/", controller.list)

	return nil
}

func (controller *ProductController) create(c *gin.Context) {
	var dto CreateDto
	c.BindJSON(&dto)

	err := controller.productService.Create(dto.Name)
	if err != nil {
		c.JSON(500, "error")
		return
	}
	c.JSON(200, "success")
}

func (controller *ProductController) list(c *gin.Context) {
	products, err := controller.productService.List()
	if err != nil {
		c.JSON(500, "error")
		return
	}
	c.JSON(200, products)
}
