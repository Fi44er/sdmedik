package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"primaryKey;type:string;" json:"id"`
	CategoryID  string `gorm:"type:uuid;not null" json:"category_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.ID = uuid.New().String()
	return nil
}
