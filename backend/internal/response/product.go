package response

type ProductFilter struct {
	Price MinMaxPrice `json:"price"`
	Count int         `json:"count"`
}

type MinMaxPrice struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
