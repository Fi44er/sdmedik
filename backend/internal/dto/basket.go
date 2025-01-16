package dto

type Create struct {
	UserID string `json:"user_id" validate:"required"`
}
