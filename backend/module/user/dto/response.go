package dto

import "github.com/Fi44er/sdmedik/backend/module/user/model"

type UserResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

func FilterUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FIO:         user.FIO,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
	}
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Count int            `json:"count"`
}
