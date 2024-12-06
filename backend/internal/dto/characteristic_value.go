package dto

type CreateCharacteristicValue struct {
	ProductID string                `json:"product_id"`
	Values    []CharacteristicValue `json:"values"`
}

type CharacteristicValue struct {
	CharacteristicID int    `json:"characteristic_id"`
	Value            string `json:"value"`
}
