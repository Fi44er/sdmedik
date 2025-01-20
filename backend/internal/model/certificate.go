package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certificate struct {
	ID              string  `gorm:"primaryKey;type:string;" json:"id"`
	CategoryArticle string  `gorm:"type:string;not null" json:"category_article"`
	RegionIso       string  `gorm:"not null" json:"region_iso"`
	Price           float64 `gorm:"not null" json:"price"`
}

func (p *Certificate) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
