package dto

type AddElement struct {
	PagePath  string `json:"page_path" validate:"required"`
	ElementID string `json:"element_id" validate:"required"`
	Value     string `json:"value" validate:"required"`
}
