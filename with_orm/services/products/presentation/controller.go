package presentation

import (
	"github.com/gin-gonic/gin"
	errorUtils "with.orm/libs/error-utils"
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
		if appError, ok := errorUtils.UnWrapWithCode(err); ok {
			c.JSON(appError.Code, appError.GetMessage())
			return
		}
	}
	c.JSON(200, "success")
}

func (controller *ProductController) list(c *gin.Context) {
	products, err := controller.productService.List()
	if err != nil {
		if appError, ok := errorUtils.UnWrapWithCode(err); ok {
			c.JSON(appError.Code, appError.GetMessage())
			return
		}
	}
	c.JSON(200, products)
}
