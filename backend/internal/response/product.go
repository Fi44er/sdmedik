package response

type ProductFilter struct {
	Price           MinMaxPrice            `json:"price"`
	Count           int                    `json:"count"`
	Characteristics []CharacteristicFilter `json:"characteristics"`
}

type CharacteristicFilter struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values []interface{} `json:"values"`
}

type MinMaxPrice struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
