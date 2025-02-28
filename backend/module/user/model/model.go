package model

type Role string

const (
	AdminRole Role = "admin"
	UserUser  Role = "user"
)

type User struct {
	ID          string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Email       string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	FIO         string `gorm:"type:varchar(255);not null" json:"fio"`
	PhoneNumber string `gorm:"type:varchar(255);not null" json:"phone_number"`
	Role        Role   `gorm:"column:role;type:role;not null;default:user" json:"role"`
	// Basket      Basket  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"basket"`
	// Orders      []Order `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"orders"`
}
