package model

type Price struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	RegionID  string  `json:"region_id"`
	Price     float64 `json:"price"`
}
