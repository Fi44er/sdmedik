package dto

type ProductResponse struct {
	ID          string   `json:"id"`
	Article     string   `json:"article"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ImageIDs    []string `json:"image_ids"`
}

type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Count    int               `json:"count"`
}
