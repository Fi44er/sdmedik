package dto

type CreateFileDTO struct {
	OwnerID   string
	OwnerType string
}

type CreateCategoryDTO struct {
	Name string `json:"name"`
}
