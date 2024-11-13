package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID              string    `gorm:"primaryKey;type:string;" json:"id"`
	FIO             string    `gorm:"type:varchar(255);not null" json:"fio"`
	Date            time.Time `gorm:"not null" json:"date"`
	Article         string    `gorm:"type:varchar(255);not null" json:"article"`
	RegionID        int       `gorm:"not null" json:"region_id"`
	PaymentMethodID int       `gorm:"not null" json:"payment_method_id"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New().String()
	return nil
}
