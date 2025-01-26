package response

type BasketResponse struct {
	ID                      string          `json:"id"`
	Quantity                int             `json:"quantity"`
	TotalPrice              float64         `json:"total_price"`
	TotalPriceWithPromotion float64         `json:"total_price_with_promotion"`
	Items                   []BasketItemRes `json:"items"`
}

type BasketItemRes struct {
	ID                      string  `json:"id"`
	Article                 string  `json:"article"`
	ProductID               string  `json:"product_id"`
	Name                    string  `json:"name"`
	Image                   string  `json:"image"`
	Quantity                int     `json:"quantity"`
	TotalPrice              float64 `json:"total_price"`
	Price                   float64 `json:"price"`
	PriceWithPromotion      float64 `json:"price_with_promotion"`
	TotalPriceWithPromotion float64 `json:"total_price_with_prom"`
	IsFree                  bool    `json:"is_free"`
}
