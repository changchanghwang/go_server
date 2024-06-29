package accountRepository

import (
	"with.framework/database"
	"with.framework/domain/account"
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

func NewAccountRepository() AccountRepository {
	return &repository{database: database.New()}
}

func (repository *repository) Save(account *account.Account) error {
	err := repository.database.Upsert(account)
	return err
}

func (repository *repository) Find() ([]*account.Account, error) {
	accounts, err := repository.database.Select(account.Account{}, nil)

	return accounts, err
}

func (repository *repository) FindOneByUserId(userId string) (*account.Account, error) {
	accounts, err := repository.database.Select(account.Account{}, database.WhereOption{UserId: userId})
	return accounts[0], err
}

func (repository *repository) Lock(userId string) {
	repository.database.Lock(database.WhereOption{UserId: userId})
}

func (repository *repository) Unlock(userId string) {
	repository.database.Unlock(database.WhereOption{UserId: userId})
}
