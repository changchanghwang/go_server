package app

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type app struct {
	*gin.Engine
}

func New() *app {
	return &app{gin.Default()}
}

func (a *app) Listen(port int) {
	a.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	a.Run(":" + strconv.Itoa(port))
}
