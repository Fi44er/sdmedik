package model

import (
	categoryModel "github.com/Fi44er/sdmedik/backend/internal/repository/category/model"
	imageModel "github.com/Fi44er/sdmedik/backend/internal/repository/image/model"
	priceModel "github.com/Fi44er/sdmedik/backend/internal/repository/price/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string                   `gorm:"primaryKey;type:string;" json:"id"`
	Name        string                   `gorm:"type:varchar(255);not null" json:"name"`
	Description string                   `gorm:"type:text" json:"description"`
	Categories  []categoryModel.Category `gorm:"many2many:product_categories;" json:"categories"`
	Prices      []priceModel.Price       `gorm:"foreignKey:ProductID" json:"prices"` // Связь с ценами
	Images      []imageModel.Image       `gorm:"foreignKey:ProductID" json:"images"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
