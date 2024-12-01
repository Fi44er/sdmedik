package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByID(ctx context.Context, id int) (model.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Category{}, err
	}

	if category.ID == 0 {
		return model.Category{}, errors.New(404, "Category not found")
	}

	return category, nil
}
