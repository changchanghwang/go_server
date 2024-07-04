package accountService

import (
	"fmt"

	"with.framework/domain/account"
	infrastructure "with.framework/infrastructure/account"
	errorUtils "with.framework/libs/error-utils"
)

type AccountService struct {
	accountRepository infrastructure.AccountRepository
}

func New(accountRepository infrastructure.AccountRepository) *AccountService {
	return &AccountService{accountRepository}
}

func (service *AccountService) AddAccount(userId string) (*account.Account, error) {
	account := account.New(userId)
	err := service.accountRepository.Save(account)
	if err != nil {
		return nil, errorUtils.WrapWithCode(err, 500)
	}

	return account, nil
}

func (service *AccountService) List() ([]*account.Account, error) {
	accounts, err := service.accountRepository.Find()
	if err != nil {
		return nil, errorUtils.WrapWithCode(err, 500)
	}

	return accounts, nil
}

func (service *AccountService) Retrieve(userId string) (*account.Account, error) {
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		return nil, errorUtils.WrapWithCode(err, 404)
	}

	return account, nil
}

func (service *AccountService) Deposit(userId string, amount int) error {
	service.accountRepository.Lock(userId)
	defer service.accountRepository.Unlock(userId)
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		return errorUtils.WrapWithCode(err, 500)
	}

	err = account.Deposit(amount)
	if err != nil {
		return errorUtils.WrapWithCode(err, 400)
	}

	err = service.accountRepository.Save(account)
	if err != nil {
		fmt.Println(err)
		//TODO: Rollback
	}
	return nil
}

func (service *AccountService) Withdraw(userId string, amount int) error {
	service.accountRepository.Lock(userId)
	defer service.accountRepository.Unlock(userId)
	account, err := service.accountRepository.FindOneByUserId(userId)
	if err != nil {
		return errorUtils.WrapWithCode(err, 400)
	}

	err = account.Withdraw(amount)
	if err != nil {
		return errorUtils.WrapWithCode(err, 400)
	}

	err = service.accountRepository.Save(account)
	if err != nil {
		// TODO: Rollback
		return errorUtils.WrapWithCode(err, 400)
	}
	return nil
}
