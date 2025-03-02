package domain

type Product struct {
	ID          string
	Article     string
	Name        string
	Description string
	Price       float64
	ImageIDs    []string
}

type ProductCategory struct {
	CategoryID string
	ProductID  string
}
