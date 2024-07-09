package infrastructure

import (
	"gorm.io/gorm"
	"with.orm/services/account/domain"
)

type AccountRepository interface {
	Save(account *domain.Account) error
	FindOne(id string) (*domain.Account, error)
}

type accountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{db}
}

func (repository *accountRepositoryImpl) Save(account *domain.Account) error {
	return repository.db.Save(account).Error
}

func (repository *accountRepositoryImpl) FindOne(id string) (*domain.Account, error) {
	account := &domain.Account{}
	err := repository.db.Where("id = ?", id).First(account).Error
	return account, err
}
