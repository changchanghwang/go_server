package router

import (
	"github.com/gofiber/fiber/v2"
	"with.framework/database"
	infra "with.framework/infrastructure/account"
	presentation "with.framework/router/account"
	application "with.framework/services/account"
)

func injectAccountController() *presentation.AccountController {
	db := database.New()
	accountRepository := infra.NewAccountRepository(db)
	accountService := application.New(accountRepository)
	accountController := presentation.New(accountService)

	return accountController
}

func Route(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	r.Route("/accounts", injectAccountController().Route)
}
