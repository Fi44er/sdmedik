package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID         string  `gorm:"primaryKey;type:string;" json:"id"`
	OrderID    string  `gorm:"type:string;not null" json:"order_id"`
	ProductID  string  `gorm:"type:string;not null" json:"product_id"`
	Name       string  `gorm:"type:string;not null" json:"name"`
	Price      float64 `gorm:"not null" json:"price"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
}

func (o *OrderItem) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
