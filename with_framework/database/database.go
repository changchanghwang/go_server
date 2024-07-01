package database

import (
	"errors"
	"sync"

	"with.framework/domain/account"
)

type accountEntity struct {
	Account *account.Account
	*sync.Mutex
}

type Database struct {
	accounts map[string]accountEntity
}

func New() *Database {
	return &Database{}
}

type WhereOption struct {
	UserId any
}

func (db *Database) Upsert(data *account.Account) error {
	if db.accounts == nil {
		db.accounts = make(map[string]accountEntity)
	}

	mutex := &sync.Mutex{}
	if existAccount, ok := db.accounts[data.Id]; ok {
		if existAccount.Account.Id != data.Id && existAccount.Account.UserId == data.UserId {
			return errors.New("already exist userId")
		}
		mutex = existAccount.Mutex
	}
	db.accounts[data.Id] = accountEntity{data, mutex}

	return nil
}

func (db *Database) Select(data account.Account, where any) ([]*account.Account, error) {
	accounts := make([]*account.Account, 0)
	whereOption, isWhereExist := where.(WhereOption)

	if isWhereExist {
		for _, account := range db.accounts {
			userId, isString := whereOption.UserId.(string)
			if isString && account.Account.UserId == userId {
				accounts = append(accounts, account.Account)
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
				accounts = append(accounts, account.Account)
			}()
		}
		wg.Wait()
		return accounts, nil
	}
}

func (db *Database) Lock(where WhereOption) {
	for _, account := range db.accounts {
		if account.Account.UserId == where.UserId {
			account.Lock()
		}
	}
}

func (db *Database) Unlock(where WhereOption) {
	for _, account := range db.accounts {
		if account.Account.UserId == where.UserId {
			account.Unlock()
		}
	}
}
