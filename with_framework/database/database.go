package database

import (
	"errors"

	"with.framework/domain/account"
)

type database struct {
	accounts map[string]account.Account
}

type WhereOption struct {
	UserId any
}

var Database = database{
	accounts: make(map[string]account.Account),
}

func (db *database) Insert(data account.Account) error {
	if _, ok := db.accounts[data.Id]; ok {
		return errors.New("account already exists")
	}
	db.accounts[data.Id] = data
	return nil
}

func (db *database) Select(data account.Account, where any) ([]account.Account, error) {
	accounts := make([]account.Account, 0)
	whereOption, whereOk := where.(WhereOption)
	for _, account := range db.accounts {
		if whereOk {
			userId, ok := whereOption.UserId.(string)
			if ok {
				if account.UserId == userId {
					accounts = append(accounts, account)
				}
			}
		} else {
			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}
