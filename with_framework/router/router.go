package router

import (
	"github.com/gofiber/fiber/v2"
	accountRouter "with.framework/router/account"
)

func Route(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	r.Route("/accounts", accountRouter.Route)
}
