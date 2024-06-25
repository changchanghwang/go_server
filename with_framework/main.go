package main

import (
	"github.com/gofiber/fiber/v2"
	"with.framework/router"
)

func main() {
	app := fiber.New()

	app.Route("/", router.Route)
	app.Listen(":3000")
}
