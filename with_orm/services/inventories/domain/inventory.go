package domain

import "with.orm/libs/entity"

type Inventory struct {
	entity.BaseModel

	Id        int    `gorm:"primaryKey"`
	ProductId string `gorm:"column:productId;unique"`
	Stock     int
}

func (Inventory) TableName() string {
	return "inventory"
}

func New(productId string, stock int) *Inventory {
	return &Inventory{
		ProductId: productId,
		Stock:     stock,
	}
}
