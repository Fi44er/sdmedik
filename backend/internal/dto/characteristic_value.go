package dto

type CharacteristicValue struct {
	CharacteristicID int       `json:"characteristic_id"`
	Value            []string  `json:"value"`
	Prices           []float64 `json:"prices"`
}
