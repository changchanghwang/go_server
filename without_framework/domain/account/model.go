package account

import "github.com/oklog/ulid/v2"

type Account struct {
	Id      string
	Balance int
	UserId  string
}

func New(userId string) *Account {
	return &Account{
		Id:      ulid.Make().String(),
		UserId:  userId,
		Balance: 0,
	}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}
