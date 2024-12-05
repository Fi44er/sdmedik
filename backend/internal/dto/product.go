package dto

type CreateProduct struct {
	Article     string `json:"article" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryIDs []int  `json:"category_ids" validate:"required"`
}
