package dto

type CreateProduct struct {
	Article              string                `json:"article" validate:"required"`
	Name                 string                `json:"name" validate:"required"`
	Description          string                `json:"description" validate:"required"`
	CategoryIDs          []int                 `json:"category_ids" validate:"required"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"required,dive"`
}

type Product struct {
	Article     string `json:"article" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
