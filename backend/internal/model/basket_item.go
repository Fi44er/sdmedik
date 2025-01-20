package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasketItem struct {
	ID         string  `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
	ProductID  string  `gorm:"type:varchar(36);not null" json:"product_id"`
	BasketID   string  `gorm:"type:varchar(36);not null" json:"basket_id"`
}

func (b *BasketItem) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()

	return nil
}
