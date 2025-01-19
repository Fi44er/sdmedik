package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasketItem struct {
	ID        string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Quantity  int    `gorm:"not null" json:"quantity"`
	ProductID string `gorm:"type:varchar(36);not null" json:"product_id"`
	BasketID  string `gorm:"type:varchar(36);not null" json:"basket_id"`
}

func (b *BasketItem) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()

	return nil
}
