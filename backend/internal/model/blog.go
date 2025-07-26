package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID      string `gorm:"primaryKey;type:varchar(36);" json:"id"`
	Preview string `gorm:"type:varchar(255);" json:"prewiew"`
	Heading string `gorm:"type:text;" json:"heading"`
	Text    string `gorm:"type:text;" json:"text"`
	Hex     string `gorm:"type:text;" json:"hex"`
}

func (m *Blog) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
