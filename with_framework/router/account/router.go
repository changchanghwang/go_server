package accountRouter

import (
	"github.com/gofiber/fiber/v2"
	errorUtils "with.framework/libs/error-utils"
	application "with.framework/services/account"
)

type CreateAccountDto struct {
	UserId string `json:"userId"`
}

type DepositDto struct {
	Amount int `json:"amount"`
}

type AccountController struct {
	accountService *application.AccountService
}

func New(accountService *application.AccountService) *AccountController {
	return &AccountController{accountService}
}

func (controller *AccountController) Route(r fiber.Router) {
	r.Post("/", func(c *fiber.Ctx) error {
		createAccountDto := CreateAccountDto{}

		if err := c.BodyParser(&createAccountDto); err != nil {
			return err
		}
		controller.accountService.AddAccount(createAccountDto.UserId)
		return nil
	})

	r.Get("/", func(c *fiber.Ctx) error {
		accounts, err := controller.accountService.List()
		if err != nil {
			if appError, ok := errorUtils.UnWrapWithCode(err); ok {
				return c.Status(appError.Code).JSON(err)
			}
		}
		return c.JSON(accounts)
	})

	/**
	 * :userId router
	 */
	userIdRouter := r.Group("/:userId")

	userIdRouter.Get("/", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		account, err := controller.accountService.Retrieve(userId)

		if err != nil {
			if appError, ok := errorUtils.UnWrapWithCode(err); ok {
				return c.Status(appError.Code).JSON(err)
			}
		}

		return c.JSON(account)
	})
	userIdRouter.Post("/deposit", func(c *fiber.Ctx) error {
		depositDto := DepositDto{}
		userId := c.Params("userId")
		if err := c.BodyParser(&depositDto); err != nil {
			return err
		}
		err := controller.accountService.Deposit(userId, depositDto.Amount)

		if err != nil {
			if appError, ok := errorUtils.UnWrapWithCode(err); ok {
				return c.Status(appError.Code).JSON(appError.Error())
			}
		}
		return nil
	})
	userIdRouter.Post("/withdraw", func(c *fiber.Ctx) error {
		depositDto := DepositDto{}
		userId := c.Params("userId")
		if err := c.BodyParser(&depositDto); err != nil {
			return err
		}
		err := controller.accountService.Withdraw(userId, depositDto.Amount)

		if err != nil {
			if appError, ok := errorUtils.UnWrapWithCode(err); ok {
				return c.Status(appError.Code).JSON(appError.Error())
			}
		}

		return nil
	})

}
