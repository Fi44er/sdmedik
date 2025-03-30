package postgres

import (
	role_model "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/role/model"
	user_model "github.com/Fi44er/sdmedik/backend/internal/module/user/infrastructure/repository/user/model"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, trigger bool, log *logger.Logger) error {

	if trigger {
		log.Info("📦 Migrating database...")
		models := []interface{}{
			role_model.Permission{},
			role_model.Role{},
			user_model.User{},
		}

		log.Info("📦 Creating types...")

		if err := db.AutoMigrate(models...); err != nil {
			log.Errorf("✖ Failed to migrate database: %v", err)
			return err
		}
	}

	log.Info("✅ Database connection successfully")
	return nil
}
