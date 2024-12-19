package image

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func (s *service) GetByProductID(ctx context.Context, productID string, tx *gorm.DB) ([]model.Image, error) {
	images, err := s.repo.GetByProductID(ctx, productID, tx)
	if err != nil {
		return nil, err
	}

	return images, nil
}
