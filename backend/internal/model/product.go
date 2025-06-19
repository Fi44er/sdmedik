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
	Price                float64               `json:"price"`
	Categories           []Category            `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE" json:"categories"`
	Images               []Image               `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"images"`
	CharacteristicValues []CharacteristicValue `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"characteristic_values"`
	BasketItems          []BasketItem          `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"basket_items"`
	Catalogs             uint8                 `gorm:"not null;default:0" json:"catalogs"`
	TRU                  string                `gorm:"type:string" json:"tru"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
