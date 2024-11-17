package model

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository/order/model"
	priceModel "github.com/Fi44er/sdmedik/backend/internal/repository/price/model"
)

type Region struct {
	ID     int                `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string             `gorm:"type:varchar(255);not null" json:"name"`
	Prices []priceModel.Price `gorm:"foreignKey:RegionID" json:"prices"`
	Orders []model.Order      `gorm:"foreignKey:PaymentMethodID" json:"orders"` // Связь с заказами
}
