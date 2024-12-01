package model

type Category struct {
	ID       int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Products []Product `gorm:"many2many:product_categories;" json:"products"`
}
