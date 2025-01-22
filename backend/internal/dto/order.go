package dto

type CreateOrder struct {
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type ChangeOrderStatus struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
