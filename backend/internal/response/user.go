package response

import "github.com/Fi44er/sdmedik/backend/internal/model"

type UserResponse struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone_number"`
	RoleID      int    `json:"role_id"`
	Role        string `json:"role"`
}

func FilterUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FIO:         user.FIO,
		PhoneNumber: user.PhoneNumber,
		RoleID:      user.RoleID,
		Role:        user.Role.Name,
	}
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Count int            `json:"count"`
}
