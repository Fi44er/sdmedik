package certificate

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) UpdateMany(ctx context.Context, data *[]model.Certificate) error {
	if err := s.repo.UpdateMany(ctx, data); err != nil {
		return err
	}
	return nil
}
