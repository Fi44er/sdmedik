package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string  `gorm:"primaryKey;type:string;" json:"id"`
	Email       string  `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string  `gorm:"type:varchar(255);not null" json:"password"`
	FIO         string  `gorm:"type:varchar(255);not null" json:"fio"`
	PhoneNumber string  `gorm:"type:varchar(255);not null" json:"phone_number"`
	RoleID      int     `gorm:"not null" json:"role_id"`
	Role        Role    `gorm:"foreignKey:RoleID" json:"role"`
	Basket      Basket  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"basket"`
	Orders      []Order `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"orders"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()

	var userRole Role
	if err := tx.First(&userRole, "name = ?", "user").Error; err != nil {
		return err
	}

	u.RoleID = userRole.ID

	return nil
}
