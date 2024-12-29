package model

type Region struct {
	ID           int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `gorm:"type:varchar(255);not null" json:"name"`
	Certificates []Certificate `gorm:"foreignKey:RegionID" json:"certificates"`
	Orders       []Order       `gorm:"foreignKey:PaymentMethodID" json:"orders"` // Связь с заказами
}
