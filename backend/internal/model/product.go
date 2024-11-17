package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string     `gorm:"primaryKey;type:string;" json:"id"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Categories  []Category `gorm:"many2many:product_categories;" json:"categories"`
	Prices      []Price    `gorm:"foreignKey:ProductID" json:"prices"` // Связь с ценами
	Images      []Image    `gorm:"foreignKey:ProductID" json:"images"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
