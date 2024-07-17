package router

import (
	"github.com/gin-gonic/gin"
	"with.orm/libs/health"
	inventoryPresentation "with.orm/services/inventories/presentation"
	productPresentation "with.orm/services/products/presentation"
)

func Route(router *gin.Engine, productController *productPresentation.ProductController, inventoryController *inventoryPresentation.InventoryController) error {
	health.Check(router)

	productRouter := router.Group("/products")
	productController.Route(productRouter)
	inventoryRouter := router.Group("/inventories")
	inventoryController.Route(inventoryRouter)

	return nil
}
