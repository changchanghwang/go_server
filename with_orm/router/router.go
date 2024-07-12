package router

import (
	"github.com/gin-gonic/gin"
	"with.orm/libs/health"
	productPresentation "with.orm/services/products/presentation"
)

func Route(router *gin.Engine, productController *productPresentation.ProductController) error {
	health.Check(router)

	productRouter := router.Group("/products")
	productController.Route(productRouter)

	return nil
}
