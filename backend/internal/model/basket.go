package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Basket struct {
	ID     string       `gorm:"primaryKey;type:string;" json:"id"`
	UserID string       `gorm:"type:string;not null" json:"user_id"`
	Items  []BasketItem `gorm:"foreignKey:BasketID;constraint:OnDelete:CASCADE" json:"items"`
}

func (b *Basket) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()

	return nil
}
