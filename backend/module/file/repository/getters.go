package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/file/converter"
	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"github.com/Fi44er/sdmedik/backend/module/file/model"
)

func (r *FileRepository) GetFilesByOwner(ctx context.Context, ownerID string, ownerType string) ([]domain.File, error) {
	var files []model.File
	if err := r.db.Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).Find(&files).Error; err != nil {
		// r.logger.Errorf()
		return nil, err
	}

	filesDomain := converter.ToDomainSliceFromModel(files)
	return filesDomain, nil
}
