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

//
// type ProductResponse struct {
// 	ID                   string                `json:"id"`
// 	Article              string                `json:"article"`
// 	Name                 string                `json:"name"`
// 	Description          string                `json:"description"`
// 	Price                float64               `json:"price"`
// 	Categories           []Category            `json:"categories"`
// 	Certificates         []Certificate         `json:"certificates"`
// 	Images               []Image               `json:"images"`
// 	CharacteristicValues []CharacteristicValue `json:"characteristic_values"`
// 	BasketItems          []BasketItem          `json:"basket_items"`
// }
