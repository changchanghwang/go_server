package app

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"with.orm/libs/health"
)

type app struct {
	router *gin.Engine
}

func New() *app {
	return &app{gin.Default()}
}

func (a *app) Listen(port int) {
	health.Check(a.router)

	a.router.Run(":" + strconv.Itoa(port))
}
