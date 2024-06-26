package account

import "github.com/google/uuid"

type Account struct {
	Id      string
	UserId  string
	Balance int
}

func New(userId string) Account {
	return Account{Id: uuid.New().String(), UserId: userId, Balance: 0}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}
