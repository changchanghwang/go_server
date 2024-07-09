package presentation

import (
	"github.com/gin-gonic/gin"
	"with.orm/services/account/application"
)

type AccountController struct {
	accountService application.AccountService
}

func NewAccountController(accountService application.AccountService) *AccountController {
	return &AccountController{accountService}
}

func (controller *AccountController) Route(router *gin.Engine) error {
	r := router.Group("/account")

	r.POST("/", controller.create)

	return nil
}

func (controller *AccountController) create(c *gin.Context) {
	userId := c.PostForm("userId")
	err := controller.accountService.Create(userId)
	if err != nil {
		c.JSON(500, "error")
		return
	}
	c.JSON(200, "success")
}
