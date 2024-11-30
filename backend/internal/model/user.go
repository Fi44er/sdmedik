package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey;type:string;" json:"id"`
	Email       string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	FIO         string `gorm:"type:varchar(255);not null" json:"fio"`
	PhoneNumber string `gorm:"type:varchar(255);not null" json:"phone_number"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
