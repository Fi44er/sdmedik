package dto

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
}
