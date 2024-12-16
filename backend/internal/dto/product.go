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

type ProductSearchCriteria struct {
	ID         string `query:"id" gorm:"id"`
	Article    string `query:"article" gorm:"article"`
	Name       string `query:"name" gorm:"name"`
	CategoryID int    `query:"category_id" gorm:"category_id"`
	Offset     int    `query:"offset" gorm:"offset"`
	Limit      int    `query:"limit" gorm:"limit"`
}
