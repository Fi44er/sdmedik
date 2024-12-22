package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID                   string                `gorm:"primaryKey;type:varchar(36);" json:"id"`
	Article              string                `gorm:"type:varchar(255);not null;unique" json:"article"`
	Name                 string                `gorm:"type:varchar(255);not null" json:"name"`
	Description          string                `gorm:"type:text" json:"description"`
	Categories           []Category            `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE" json:"categories"`
	Prices               []Price               `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"prices"` // Связь с ценами
	Images               []Image               `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"images"`
	CharacteristicValues []CharacteristicValue `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"characteristic_values"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
