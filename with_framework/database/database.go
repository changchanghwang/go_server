package database

import (
	"errors"
	"sync"

	"with.framework/domain/account"
	errorUtils "with.framework/libs/error-utils"
)

type accountEntity struct {
	Account *account.Account
	*sync.Mutex
}

type Database struct {
	accounts sync.Map
}

func New() *Database {
	return &Database{}
}

type WhereOption struct {
	UserId any
}

func (db *Database) Upsert(data *account.Account) error {
	mutex := &sync.Mutex{}

	if existAccount, isExist := db.accounts.Load(data.Id); isExist {
		account := existAccount.(accountEntity)
		mutex = account.Mutex
	} else {
		isUniqueExist := false

		db.accounts.Range(func(key, value interface{}) bool {
			account := value.(accountEntity)

			if account.Account.UserId == data.UserId {
				isUniqueExist = true
				return false
			}

			return true
		})

		if isUniqueExist {
			return errorUtils.Wrap(errors.New("already exist userId"))
		}
	}

	db.accounts.Store(data.Id, accountEntity{Account: data, Mutex: mutex})
	return nil
}

func (db *Database) Select(data account.Account, where any) ([]*account.Account, error) {
	accounts := make([]*account.Account, 0)
	whereOption, isWhereExist := where.(WhereOption)
	db.accounts.Range(func(key, value interface{}) bool {
		account := value.(accountEntity)
		if isWhereExist {
			userId, isString := whereOption.UserId.(string)
			if isString && account.Account.UserId == userId {
				accounts = append(accounts, account.Account)
			}
		} else {
			accounts = append(accounts, account.Account)
		}
		return true
	})
	return accounts, nil
}

func (db *Database) Lock(where WhereOption) {
	db.accounts.Range(func(key, value interface{}) bool {
		accountEntity := value.(accountEntity)
		if accountEntity.Account.UserId == where.UserId {
			accountEntity.Lock()
			return false
		}
		return true
	})
}

func (db *Database) Unlock(where WhereOption) {
	db.accounts.Range(func(key, value interface{}) bool {
		accountEntity := value.(accountEntity)
		if accountEntity.Account.UserId == where.UserId {
			accountEntity.Unlock()
			return false
		}
		return true
	})
}
