package dto

type CreateProduct struct {
	Article              string                `json:"article" validate:"required"`
	Name                 string                `json:"name" validate:"required" msg:"Name is required"`
	Description          string                `json:"description" validate:"required"`
	CategoryIDs          []int                 `json:"category_ids"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"dive"`
}

type Product struct {
	Article     string `json:"article" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
