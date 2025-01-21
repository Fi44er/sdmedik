package response

type BasketResponse struct {
	ID         string          `json:"id"`
	Quantity   int             `json:"quantity"`
	TotalPrice float64         `json:"total_price"`
	Items      []BasketItemRes `json:"items"`
}

type BasketItemRes struct {
	ID         string  `json:"id"`
	Article    string  `json:"article"`
	ProductID  string  `json:"product_id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Price      float64 `json:"price"`
}
