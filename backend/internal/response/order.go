package response

type OrderResponse struct {
	ID         string              `json:"id"`
	Email      string              `json:"email"`
	Phone      string              `json:"phone"`
	FIO        string              `json:"fio"`
	Status     string              `json:"status"`
	TotalPrice float64             `json:"total_price"`
	Items      []OrderItemResponse `json:"items"`
	CreatedAt  string              `json:"created_at"`
}

type OrderItemResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
