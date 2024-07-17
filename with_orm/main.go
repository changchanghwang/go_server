package main

import (
	"go.uber.org/fx"
	"with.orm/app"
	"with.orm/libs/db"

	inventoryApplication "with.orm/services/inventories/application"
	inventoryInfrastructure "with.orm/services/inventories/infrastructure"
	inventoryPresentation "with.orm/services/inventories/presentation"

	productApplication "with.orm/services/products/application"
	productInfrastructure "with.orm/services/products/infrastructure"
	productPresentation "with.orm/services/products/presentation"
)

func main() {
	fx.New(
		fx.Provide(db.Init),
		fx.Provide(productPresentation.NewProductController),
		fx.Provide(productApplication.NewProductService),
		fx.Provide(productInfrastructure.NewProductRepository),
		fx.Provide(inventoryPresentation.NewInventoryController),
		fx.Provide(inventoryApplication.NewInventoryService),
		fx.Provide(inventoryInfrastructure.NewInventoryRepository),
		fx.Invoke(app.ServerRun),
	).Run()
}
