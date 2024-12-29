package dto

type CreateProduct struct {
	Article              string                `json:"article" validate:"required"`
	Name                 string                `json:"name" validate:"required"`
	Description          string                `json:"description" validate:"required"`
	Price                float64               `json:"price"`
	CategoryIDs          []int                 `json:"category_ids"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"dive"`
}

type Product struct {
	Article     string  `json:"article" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price"`
}

type ProductSearchCriteria struct {
	ID         string `query:"id" gorm:"id"`
	Article    string `query:"article" gorm:"article"`
	Name       string `query:"name" gorm:"name"`
	CategoryID int    `query:"category_id" gorm:"category_id"`
	Offset     int    `query:"offset" gorm:"offset"`
	Limit      int    `query:"limit" gorm:"limit"`
}

type UpdateProduct struct {
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	Price                float64               `json:"price"`
	DelImages            []DelImage            `json:"del_images" validate:"dive"`
	CategoryIDs          []int                 `json:"category_ids"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"dive"`
}

type DelImage struct {
	Name string `json:"name" validate:"required"`
	ID   string `json:"id" validate:"required"`
}
