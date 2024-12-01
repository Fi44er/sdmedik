package dto

type Register struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	FIO         string `json:"fio" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerifyCode struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type SendCode struct {
	Email string `json:"email" validate:"required,email"`
}
