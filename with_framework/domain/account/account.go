package account

import (
	"errors"

	"github.com/google/uuid"
)

type Account struct {
	Id      string
	UserId  string
	Balance int
}

func New(userId string) *Account {
	return &Account{Id: uuid.New().String(), UserId: userId, Balance: 0}
}

func (a *Account) Deposit(amount int) error {
	if amount < 0 {
		return errors.New(("amount must be positive"))
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("amount must be positive")
	}
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount
	return nil
}
