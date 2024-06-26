package database

import (
	"errors"
	"sync"

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
	whereOption, isWhereExist := where.(WhereOption)

	if isWhereExist {
		for _, account := range db.accounts {
			userId, isString := whereOption.UserId.(string)
			if isString && account.UserId == userId {
				accounts = append(accounts, account)
			}
		}
		return accounts, nil
	} else {
		// where절이 없으면 비동기로 map을 돌면서 slice로 만들어서 반환한다.
		var wg sync.WaitGroup
		wg.Add(len(db.accounts))

		for _, account := range db.accounts {
			go func() {
				defer wg.Done()
				accounts = append(accounts, account)
			}()
		}
		wg.Wait()
		return accounts, nil
	}
}
