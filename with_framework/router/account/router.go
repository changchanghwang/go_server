package accountRouter

import (
	"github.com/gofiber/fiber/v2"
	accountService "with.framework/services/account"
)

type CreateAccountDto struct {
	UserId string `json:"userId"`
}

func Route(r fiber.Router) {

	r.Post("/", func(c *fiber.Ctx) error {
		createAccountDto := CreateAccountDto{}

		if err := c.BodyParser(&createAccountDto); err != nil {
			return err
		}
		accountService.AddAccount(createAccountDto.UserId)
		return nil
	})

	r.Get("/", func(c *fiber.Ctx) error {
		accounts := accountService.List()
		return c.JSON(accounts)
	})

	r.Get("/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		account := accountService.Retrieve(userId)
		return c.JSON(account)
	})

}