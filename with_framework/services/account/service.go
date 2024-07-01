package accountService

import (
	"fmt"

	"with.framework/domain/account"
	infrastructure "with.framework/infrastructure/account"
)

type AccountService struct {
	accountRepository infrastructure.AccountRepository
}

func New(accountRepository infrastructure.AccountRepository) *AccountService {
	return &AccountService{accountRepository}
}

func (service *AccountService) AddAccount(userId string) *account.Account {
	account := account.New(userId)
	err := service.accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
	}

	return account
}

func (service *AccountService) List() []*account.Account {
	accounts, err := service.accountRepository.Find()
	if err != nil {
		fmt.Println(err)
	}

	return accounts
}

func (service *AccountService) Retrieve(userId string) *account.Account {
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		fmt.Println(err)
	}

	return account
}

func (service *AccountService) Deposit(userId string, amount int) {
	service.accountRepository.Lock(userId)
	defer service.accountRepository.Unlock(userId)
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = account.Deposit(amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = service.accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
		//TODO: Rollback
	}
}

func (service *AccountService) Withdraw(userId string, amount int) {
	service.accountRepository.Lock(userId)
	defer service.accountRepository.Unlock(userId)
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = account.Withdraw(amount)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = service.accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
		// TODO: Rollback
	}
}
