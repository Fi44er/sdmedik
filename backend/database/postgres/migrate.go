package database

import (
	category_model "github.com/Fi44er/sdmedik/backend/module/category/model"
	file_model "github.com/Fi44er/sdmedik/backend/module/file/model"
	prduct_model "github.com/Fi44er/sdmedik/backend/module/product/model"
	user_model "github.com/Fi44er/sdmedik/backend/module/user/model"
	"github.com/Fi44er/sdmedik/backend/shared/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			&user_model.User{},
			&file_model.File{},
			&category_model.Category{},
			&prduct_model.Product{},
			&prduct_model.ProductCategory{},
		}

		log.Info("📦 Creating types...")
		db.Exec("CREATE TYPE role AS ENUM('admin', 'user')")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}
