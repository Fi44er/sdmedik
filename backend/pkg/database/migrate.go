package database

import (
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&model.User{},
		&model.Role{},
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

	if err := createDefaultRole(db); err != nil {
		return err
	}

	return nil
}

func createDefaultRole(db *gorm.DB) error {
	roles := []model.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
