package model

type Region struct {
	ID           int           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `gorm:"type:varchar(255);not null;unique" json:"name"`
	Iso3166      string        `gorm:"type:varchar(255);not null;unique" json:"iso3166"`
	Markup       float64       `json:"markup"`
	Certificates []Certificate `gorm:"foreignKey:RegionIso;references:Iso3166" json:"certificates"`
}
