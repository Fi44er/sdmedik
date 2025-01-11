package model

type OrderItem struct {
	ID        string  `gorm:"primaryKey;type:string;" json:"id"`
	OrderID   string  `gorm:"type:string;not null" json:"order_id"`
	ProductID string  `gorm:"type:string;not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}
