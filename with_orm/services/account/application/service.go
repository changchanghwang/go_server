package application

import (
	account "with.orm/services/account/domain"
	"with.orm/services/account/infrastructure"
)

type AccountService struct {
	accountRepository infrastructure.AccountRepository
}

func NewAccountService(accountRepository infrastructure.AccountRepository) *AccountService {
	return &AccountService{accountRepository}
}

func (service *AccountService) Create(userId string) error {
	account := account.New(userId)
	return service.accountRepository.Save(account)
}
