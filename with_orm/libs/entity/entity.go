package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time `json:"_", gorm:"autoCreateTime:nano;"`
	UpdatedAt time.Time `json:"__", gorm:"autoUpdateTime:nano;"`
}

type SoftDeletableModel struct {
	BaseModel
	DeletedAt gorm.DeletedAt `json:"_", gorm:"index;"`
}
