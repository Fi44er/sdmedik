package dto

type CreateCharacteristic struct {
	Name       string `json:"name" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
	DataType   string `json:"data_type" validate:"required;characteristic_type"`
}

type CharacteristicWithoutCategoryID struct {
	Name     string `json:"name" validate:"required"`
	DataType string `json:"data_type" validate:"characteristic_type"`
}
