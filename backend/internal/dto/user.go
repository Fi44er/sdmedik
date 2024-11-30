package dto

type UpdateUser struct {
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
