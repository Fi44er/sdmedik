package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetAll(ctx context.Context) (*[]model.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(*categories) == 0 {
		return nil, errors.New(404, "Categories not found")
	}

	return categories, nil
}
