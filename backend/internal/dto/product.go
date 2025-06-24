package dto

type CreateProduct struct {
	Article              string                `json:"article" validate:"required"`
	Name                 string                `json:"name" validate:"required"`
	Description          string                `json:"description"`
	Price                float64               `json:"price"`
	CategoryIDs          []int                 `json:"category_ids"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"dive"`
	TRU                  string                `json:"tru"`
}

type Product struct {
	Article     string  `json:"article" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price"`
	TRU         string  `json:"tru"`
}

type ProductSearchCriteria struct {
	ID         string         `query:"id" gorm:"id"`
	Article    string         `query:"article" gorm:"article"`
	Name       string         `query:"name" gorm:"name"`
	CategoryID int            `query:"category_id" gorm:"category_id"`
	Offset     int            `query:"offset" gorm:"offset"`
	Limit      int            `query:"limit" gorm:"limit"`
	Filters    ProductFilters `query:"filters" gorm:"-"`
	Minimal    bool           `query:"minimal" gorm:"-"`
	Catalogs   []int          `query:"catalogs" gorm:"-"`
	Iso        string         `query:"iso" gorm:"-"`
}

type ProductFilters struct {
	Price           PriceFilter            `json:"price"`
	Characteristics []FilterCharacteristic `json:"characteristics" validate:"dive"`
}

type FilterCharacteristic struct {
	CharacteristicID int      `json:"characteristic_id" validate:"required"`
	Values           []string `json:"values" validate:"required"`
}

type PriceFilter struct {
	Min float64 `json:"min" example:"20"`
	Max float64 `json:"max" example:"100"`
}

type UpdateProduct struct {
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	Price                float64               `json:"price"`
	DelImages            []DelImage            `json:"del_images" validate:"dive"`
	CategoryIDs          []int                 `json:"category_ids"`
	CharacteristicValues []CharacteristicValue `json:"characteristic_values" validate:"dive"`
	Catalogs             []int                 `json:"catalogs"`
	TRU                  string                `json:"tru"`
}

type DelImage struct {
	Name string `json:"name" validate:"required"`
	ID   string `json:"id" validate:"required"`
}

type ProductFilter struct {
	Price           MinMaxPrice            `json:"price" validate:"dive"`
	Characteristics []CharacteristicFilter `json:"characteristics"`
}

type CharacteristicFilter struct {
	Name   string        `json:"name"`
	Values []interface{} `json:"values"`
}

type MinMaxPrice struct {
	Min float64 `json:"min" validate:"required"`
	Max float64 `json:"max" validate:"required"`
}
