package dto

type CategoryResponse struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	ImageIDs []string `json:"image_ids"`
}
