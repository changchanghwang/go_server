package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	Id      string `gorm: "primaryKey"`
	UserId  string `gorm: "uniqueIndex"`
	Balance int
}

func (Account) TableName() string {
	return "account"
}

func New(userId string) *Account {
	return &Account{Id: uuid.New().String(), UserId: userId, Balance: 0}
}
