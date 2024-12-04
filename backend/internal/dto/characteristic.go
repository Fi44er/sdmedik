package dto

type CreateCharacteristic struct {
	Name       string `json:"name" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
}
