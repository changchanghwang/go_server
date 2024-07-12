package main

import (
	"go.uber.org/fx"
	"with.orm/app"
	"with.orm/libs/db"
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
		fx.Invoke(app.ServerRun),
	).Run()
}
