package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/file/converter"
	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"gorm.io/gorm"
)

func (r *FileRepository) CreateMany(ctx context.Context, filesDomains []domain.File, tx *gorm.DB) error {
	r.logger.Info("creating files...")
	db := tx
	if db == nil {
		db = r.db
	}
	filesModels := converter.ToModelSliceFromDomain(filesDomains)

	// if len(filesDomains) > 0 { // Условие можно настроить под нужды
	// 	errMsg := "intentional error to test transaction rollback"
	// 	r.logger.Errorf("Artificial error: %v", errMsg)
	// 	return fmt.Errorf(errMsg)
	// }

	if err := db.WithContext(ctx).Create(filesModels).Error; err != nil {
		r.logger.Errorf("error creating files: %v", err)
		return err
	}
	r.logger.Info("files created")
	return nil
}
