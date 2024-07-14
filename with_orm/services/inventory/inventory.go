package domain

import "with.orm/libs/entity"

type Inventory struct {
	entity.BaseModel

	Id        int    `gorm: "primaryKey"`
	productId string `gorm: "unique"`
	stock     int
}

func (Inventory) TableName() string {
	return "inventory"
}

func New(productId string, stock int) *Inventory {
	return &Inventory{
		productId: productId,
		stock:     stock,
	}
}
