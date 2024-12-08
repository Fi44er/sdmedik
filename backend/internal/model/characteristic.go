package model

type Type string

const (
	TypeString Type = "string"
	TypeInt    Type = "int"
	TypeFloat  Type = "float"
	TypeBool   Type = "bool"
)

type Characteristic struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID int    `gorm:"not null" json:"category_id"`
	DataType   Type   `gorm:"type:varchar(50);not null" json:"data_type"`
}
