package model

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository/tokens/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey;type:string;" json:"id"`
	Login    string `gorm:"type:varchar(100);unique;not null" json:"login"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`

	Tokens []model.Token `gorm:"foreignKey:UserID" json:"tokens"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
