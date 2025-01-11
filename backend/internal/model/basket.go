package model

type Basket struct {
	ID     string       `gorm:"primaryKey;type:string;" json:"id"`
	UserID string       `gorm:"type:string;not null" json:"user_id"`
	Items  []BasketItem `gorm:"foreignKey:BasketID" json:"items"`
}
