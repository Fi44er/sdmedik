package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Price struct {
	ID        string  `gorm:"primaryKey;type:string;" json:"id"`
	ProductID string  `gorm:"type:string;not null" json:"product_id"`
	RegionID  string  `gorm:"type:string;not null" json:"region_id"`
	Price     float64 `gorm:"not null" json:"price"`
}

func (p *Price) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
