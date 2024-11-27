package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey;type:string;" json:"id"`
	Email       string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	FIO         string `gorm:"type:varchar(255)" json:"fio"`
	PhoneNumber string `gorm:"type:varchar(255)" json:"phone_number"`

	Tokens []Token `gorm:"foreignKey:UserID" json:"tokens"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
