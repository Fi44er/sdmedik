package model

type Category struct {
	ID              int              `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string           `gorm:"type:varchar(255);not null;unique" json:"name"`
	Products        []Product        `gorm:"many2many:product_categories;constraint:onDelete:CASCADE" json:"products"`
	Characteristics []Characteristic `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"characteristic"`
	Images          []Image          `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"images"`
}
