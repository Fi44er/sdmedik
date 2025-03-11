package dto

import "github.com/gofiber/fiber/v2/middleware/session"

type CreateBasket struct {
	UserID string `json:"user_id" validate:"required"`
}

type AddBasketItem struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type UpdateItemQuantity struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type MoveBasket struct {
	UserID  string
	Session *session.Session
}
