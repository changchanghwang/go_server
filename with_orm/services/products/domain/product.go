package domain

import (
	"github.com/google/uuid"
	"with.orm/libs/entity"
)

type Product struct {
	entity.BaseModel `gorm:"embedded"`

	Id   string `gorm: "primaryKey"`
	Name string
}

func (Product) TableName() string {
	return "product"
}

func New(name string) *Product {
	return &Product{
		Id:   uuid.New().String(),
		Name: name,
	}
}
