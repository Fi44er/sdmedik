package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID           string         `gorm:"primaryKey;type:string;" json:"id"`
	UserID       sql.NullString `gorm:"type:string;" json:"user_id"`
	Email        string         `gorm:"type:string;not null" json:"email"`
	Phone        string         `gorm:"type:string;not null" json:"phone"`
	FIO          string         `gorm:"type:string;not null" json:"fio"`
	Address      string         `gorm:"type:string;not null" json:"address"`
	TotalPrice   float64        `gorm:"not null" json:"total_price"`
	Status       string         `gorm:"column:status;type:status;not null" json:"status"` // pending or completed
	Items        []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	FragmentLink string         `gorm:"type:string" json:"fragment_link"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID = uuid.New().String()
	return nil
}
