package dto

type UserDTO struct {
	FIO         string `json:"fio" validate:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
