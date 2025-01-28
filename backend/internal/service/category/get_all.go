package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) GetAll(ctx context.Context) (*[]model.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(*categories) == 0 {
		return nil, constants.ErrCategoryNotFound
	}

	return categories, nil
}
