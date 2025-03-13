package response

type ProductFilter struct {
	Price           MinMaxPrice            `json:"price"`
	Count           int                    `json:"count"`
	Characteristics []CharacteristicFilter `json:"characteristics"`
}

type CharacteristicFilter struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
	Values []interface{} `json:"values"`
}

type MinMaxPrice struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type ProductResponse struct {
	ID               string                     `json:"id"`
	Article          string                     `json:"article"`
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Price            float64                    `json:"price"`
	CertificatePrice float64                    `json:"certificate_price"`
	Categories       []ProductCategoryRes       `json:"categories"`
	Images           []ProductImageRes          `json:"images"`
	Characteristic   []ProductCharacteristicRes `json:"characteristic"`
	Catalogs         uint8                      `json:"catalogs"`
}

type ProductCategoryRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductImageRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductCharacteristicRes struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
	Name  string `json:"name"`
}

type ProductPopularity struct {
	ProductID  string `gorm:"column:product_id" json:"product_id"`
	OrderCount int    `gorm:"column:order_count" json:"order_count"`
}

type TopProductRes struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	OrderCount int     `json:"order_count"`
	Article    string  `json:"article"`
	Image      string  `json:"image"`
	Name       string  `json:"name"`
}
