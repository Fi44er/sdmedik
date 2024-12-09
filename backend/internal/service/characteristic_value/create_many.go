package characteristicvalue

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func (s *service) CreateMany(ctx context.Context, characteristicValues *[]model.CharacteristicValue, tx *gorm.DB) error {
	if err := s.repo.CreateMany(ctx, characteristicValues, tx); err != nil {
		return err
	}
	return nil
}
