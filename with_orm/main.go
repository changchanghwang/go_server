package main

import (
	"with.orm/app"
)

func main() {
	app := app.New()
	app.Listen(3000)
}
