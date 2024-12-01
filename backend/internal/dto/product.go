package dto

type CreateProduct struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryIDs []int  `json:"category_ids" validate:"required"`
}
