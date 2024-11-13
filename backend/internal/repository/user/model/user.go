package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;type:string;" json:"id"`
	Login    string `gorm:"type:varchar(100);unique;not null" json:"login"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
