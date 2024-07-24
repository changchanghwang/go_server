package presentation

import (
	"github.com/gin-gonic/gin"
	errorUtils "with.orm/libs/error-utils"
	"with.orm/services/inventories/application"
)

type InventoryController struct {
	inventoryService *application.InventoryService
}

func NewInventoryController(inventoryService *application.InventoryService) *InventoryController {
	return &InventoryController{inventoryService}
}

func (controller *InventoryController) Route(router *gin.RouterGroup) error {
	router.POST("/", controller.create)
	router.GET("/", controller.list)

	return nil
}

func (controller *InventoryController) create(c *gin.Context) {
	var dto CreateDto
	c.BindJSON(&dto)

	err := controller.inventoryService.Create(dto.ProductId, dto.Stock)
	if err != nil {
		if appError, ok := errorUtils.UnWrapWithCode(err); ok {
			c.JSON(appError.Code, appError.GetMessage())
			return
		}
	}
	c.JSON(200, "success")
}

func (controller *InventoryController) list(c *gin.Context) {
	inventories, err := controller.inventoryService.List()
	if err != nil {
		if appError, ok := errorUtils.UnWrapWithCode(err); ok {
			c.JSON(appError.Code, appError.GetMessage())
			return
		}
	}
	c.JSON(200, inventories)
}
