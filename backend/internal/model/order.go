package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID              string      `gorm:"primaryKey;type:string;" json:"id"`
	UserID          string      `gorm:"type:string;not null" json:"user_id"`
	CartID          *string     `gorm:"type:string;" json:"cart_id"`
	PaymentMethodID int         `gorm:"not null" json:"payment_method_id"`
	TotalAmount     float64     `gorm:"not null" json:"total_amount"`
	Status          string      `gorm:"not null" json:"status"` // pending or completed
	Items           []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt       time.Time   `gorm:"autoCreateTime" json:"created_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
