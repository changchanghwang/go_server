package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"with.orm/libs/db"
	"with.orm/router"
	inventoryPresentation "with.orm/services/inventories/presentation"
	productPresentation "with.orm/services/products/presentation"
)

type App struct {
	router *gin.Engine
	server *http.Server
}

func new() *App {
	router := gin.Default()
	return &App{router, &http.Server{Handler: router.Handler()}}
}

func (a *App) Start(port int) error {
	a.server.Addr = ":" + strconv.Itoa(port)
	return a.server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func ServerRun(lc fx.Lifecycle, productController *productPresentation.ProductController, inventoryController *inventoryPresentation.InventoryController) *App {
	server := new()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			db.Init()
			router.Route(server.router, productController, inventoryController)
			go func() {
				// service connections
				fmt.Println("Server running on 3000")
				server.Start(3000)
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Println("Server exiting")
			server.Stop(c)
			return nil
		},
	})
	return server
}
