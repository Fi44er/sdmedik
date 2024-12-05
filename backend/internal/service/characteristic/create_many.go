package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) CreateMany(ctx context.Context, characteristics *[]model.Characteristic) error {
	if err := s.repo.CreateMany(ctx, characteristics); err != nil {
		return err
	}
	return nil
}
