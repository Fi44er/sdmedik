package model

type BasketItem struct {
	ID        string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Quantity  int    `gorm:"not null" json:"quantity"`
	ProductID string `gorm:"type:varchar(36);not null" json:"product_id"`
	BasketID  string `gorm:"type:varchar(36);not null" json:"basket_id"`
}
