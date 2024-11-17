package database

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&model.User{},
		&model.Token{},
		&model.Region{},
		&model.ProductCategory{},
		&model.Product{},
		&model.Category{},
		&model.Image{},
		&model.Price{},
		&model.Order{},
		&model.PaymentMethod{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
