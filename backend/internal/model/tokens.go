package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID       string    `gorm:"primaryKey;type:string;" json:"id"`
	Token    string    `gorm:"type:text;not null" json:"token"`
	UserID   string    `gorm:"type:string;not null" json:"user_id"`
	CreateAt time.Time `gorm:"autoCreateTime" json:"create_at"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New().String()
	return nil
}
