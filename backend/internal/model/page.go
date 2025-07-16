package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	Path string `json:"path" gorm:"primaryKey;type:string;not null"`

	Elements []Element `json:"elements" gorm:"foreignKey:PagePath;constraint:OnDelete:CASCADE;"`
}

type Element struct {
	ID        string `json:"id" gorm:"primaryKey;type:string;"`
	ElementID string `json:"element_id" gorm:"type:string;not null;uniqueIndex:idx_element_page"`
	Value     string `json:"value" gorm:"type:text;not null"`
	PagePath  string `json:"page_path" gorm:"type:string;not null;uniqueIndex:idx_element_page"`
}

func (e *Element) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New().String()
	return nil
}
