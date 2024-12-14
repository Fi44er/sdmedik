package category

import (
	"context"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByIDs(ctx context.Context, ids []int) ([]model.Category, error) {
	categories, err := s.repo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
