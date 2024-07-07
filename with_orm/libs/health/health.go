package health

import (
	"github.com/gin-gonic/gin"
)

func Check(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		//TODO: DB연결도 체크?
		c.JSON(200, "pong")
	})
}
