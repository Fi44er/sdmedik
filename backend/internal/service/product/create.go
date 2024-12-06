package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, product *dto.CreateProduct) error {
	s.logger.Info("Creating product in service...")

	if err := s.validator.Struct(product); err != nil {
		return errors.New(400, err.Error())
	}

	var modelProduct model.Product
	if err := utils.DtoToModel(product, &modelProduct); err != nil {
		return err
	}

	categories, err := s.categoryService.GetByIDs(ctx, product.CategoryIDs)
	if err != nil {
		return err
	}

	modelProduct.Categories = categories

	if err := s.repo.Create(ctx, &modelProduct); err != nil {
		return err
	}

	s.logger.Info("Product created successfully")
	return nil
}
