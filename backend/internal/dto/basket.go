package dto

import "github.com/gofiber/fiber/v2/middleware/session"

type CreateBasket struct {
	UserID string `json:"user_id" validate:"required"`
}

type AddBasketItem struct {
	Iso       string `json:"iso"`
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`

	DinamicOptions []DinamicOption `json:"dynamic_options"` // выбранные характеристики
}

type DinamicOption struct {
	ID    int    `json:"id" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type UpdateItemQuantity struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type MoveBasket struct {
	UserID  string
	Session *session.Session
}
