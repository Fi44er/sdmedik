package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	ID        string `gorm:"primaryKey;type:string;" json:"id"`
	ProductID string `gorm:"type:uuid;not null" json:"product_id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
}

func (i *Image) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.New().String()
	return nil
}
