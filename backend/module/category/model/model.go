package model

type Category struct {
	ID       string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Name     string   `json:"name"`
	ImageIDs []string `gorm:"-" json:"image_ids"`
}
