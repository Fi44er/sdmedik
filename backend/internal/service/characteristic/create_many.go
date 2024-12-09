package characteristic

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"gorm.io/gorm"
)

func (s *service) CreateMany(ctx context.Context, characteristics *[]model.Characteristic, tx *gorm.DB) error {
	if err := s.repo.CreateMany(ctx, characteristics, tx); err != nil {
		return err
	}
	return fmt.Errorf("Characteristics created successfully")
}
