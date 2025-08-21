package model

type CharacteristicValue struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Value            []string  `gorm:"type:json;serializer:json;not null" json:"value"`
	CharacteristicID int       `gorm:"not null" json:"characteristic_id"`
	ProductID        string    `gorm:"not null" json:"product_id"`
	Prices           []float64 `gorm:"type:json;serializer:json;" json:"price"`
}
