package model

type Characteristic struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID int    `gorm:"not null" json:"category_id"`
}