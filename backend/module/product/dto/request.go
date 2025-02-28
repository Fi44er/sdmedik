package dto

type CreateProductDTO struct {
	Article     string `json:"article" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}

type CreateFileDTO struct {
	OwnerID   string
	OwnerType string
}
