package service

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	customerr "github.com/Fi44er/sdmedik/backend/shared/custom_err"
)

func (s *CategoryService) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		s.logger.Infof("Category with id %s not found", id)
		return nil, customerr.ErrCategoryNotFound
	}

	return category, nil
}

func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		s.logger.Info("Categories not found")
		return nil, customerr.ErrCategoryNotFound
	}

	return categories, nil
}
