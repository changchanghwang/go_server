package accountService

import (
	"fmt"

	"with.framework/domain/account"
	accountRepository "with.framework/infrastructure/account"
)

func AddAccount(userId string) {
	account := account.New(userId)

	err := accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
	}
}

func List() []account.Account {
	accounts, err := accountRepository.Find()
	if err != nil {
		fmt.Println(err)
	}

	return accounts
}

func Retrieve(userId string) account.Account {
	account, err := accountRepository.FindOneByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}

	return account
}

func Deposit(userId string, amount int) {
	account, err := accountRepository.FindOneByUserId(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = account.Deposit(amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
		//TODO: Rollback
	}
}
