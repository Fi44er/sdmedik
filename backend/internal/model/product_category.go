package model

type ProductCategory struct {
	ProductID  string `gorm:"primaryKey" json:"product_id"`
	CategoryID string `gorm:"primaryKey" json:"category_id"`
}
