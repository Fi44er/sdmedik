package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) GetByID(ctx context.Context, id int) (*model.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, constants.ErrCategoryNotFound
	}

	return category, nil
}
