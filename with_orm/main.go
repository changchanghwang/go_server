package main

import (
	"with.orm/app"
	"with.orm/libs/db"
)

func main() {
	db.Init()
	app := app.New()
	app.Listen(3000)
}
