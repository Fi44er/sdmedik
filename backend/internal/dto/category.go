package dto

type CreateCategory struct {
	Name            string   `json:"name" validate:"required"`
	Characteristics []string `json:"characteristics"`
}

type CategoryWithoutCharacteristics struct {
	Name string `json:"name"`
}
