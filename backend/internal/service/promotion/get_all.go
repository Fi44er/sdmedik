package promotion

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetAll(ctx context.Context) (*[]model.Promotion, error) {
	promotions, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return promotions, nil
}
