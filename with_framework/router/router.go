package router

import "github.com/gofiber/fiber/v2"

func Route(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
