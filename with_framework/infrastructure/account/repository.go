package accountRepository

import (
	"with.framework/database"
	"with.framework/domain/account"
)

func Save(account account.Account) error {
	err := database.Database.Insert(account)
	return err
}

func Find() ([]account.Account, error) {
	accounts, err := database.Database.Select(account.Account{}, nil)

	return accounts, err
}

func FindOneByUserId(userId string) (account.Account, error) {
	accounts, err := database.Database.Select(account.Account{}, database.WhereOption{UserId: userId})
	return accounts[0], err
}
