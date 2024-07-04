package accountRepository

import (
	"with.framework/database"
	"with.framework/domain/account"
	errorUtils "with.framework/libs/error-utils"
)

type AccountRepository interface {
	Save(account *account.Account) error
	Find() ([]*account.Account, error)
	FindOneByUserId(userId string) (*account.Account, error)
	Lock(userId string)
	Unlock(userId string)
}

type repository struct {
	database *database.Database
}

func NewAccountRepository(database *database.Database) AccountRepository {
	return &repository{database}
}

func (repository *repository) Save(account *account.Account) error {
	err := repository.database.Upsert(account)
	if err != nil {
		return errorUtils.Wrap(err)
	}
	return nil
}

func (repository *repository) Find() ([]*account.Account, error) {
	accounts, err := repository.database.Select(account.Account{}, nil)
	if err != nil {
		return nil, errorUtils.Wrap(err)
	}

	return accounts, nil
}

func (repository *repository) FindOneByUserId(userId string) (*account.Account, error) {
	accounts, err := repository.database.Select(account.Account{}, database.WhereOption{UserId: userId})
	if len(accounts) == 0 {
		return nil, errorUtils.Wrap(err)
	}
	if err != nil {
		return nil, errorUtils.Wrap(err)
	}
	return accounts[0], nil
}

func (repository *repository) Lock(userId string) {
	repository.database.Lock(database.WhereOption{UserId: userId})
}

func (repository *repository) Unlock(userId string) {
	repository.database.Unlock(database.WhereOption{UserId: userId})
}
