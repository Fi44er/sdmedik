package dto

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
