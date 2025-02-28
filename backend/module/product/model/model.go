package model

type Product struct {
	ID          string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Article     string   `gorm:"not null" json:"article"`
	Name        string   `gorm:"not null" json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ImageIDs    []string `gorm:"-" json:"image_ids"`
}
