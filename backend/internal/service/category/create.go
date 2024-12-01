package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, category *dto.CreateCategory) error {
	s.logger.Info("Creating category in service...")

	if err := s.validator.Struct(category); err != nil {
		return errors.New(400, err.Error())
	}

	var modelCategory model.Category
	if err := utils.DtoToModel(category, &modelCategory); err != nil {
		return err
	}
	if err := s.repo.Create(ctx, &modelCategory); err != nil {
		return err
	}
	s.logger.Info("Category created successfully")
	return nil
}
