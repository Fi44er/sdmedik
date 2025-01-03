package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	// "github.com/Fi44er/sdmedik/backend/internal/response"
	// "github.com/samber/lo"
)

func (s *service) GetFilter(ctx context.Context, categoryID int) error {
	// filterRes := new(response.ProductFilter)
	category, err := s.categoryService.GetByID(ctx, categoryID)
	if err != nil {
		s.logger.Errorf("Error getting category: %v", err)
		return err
	}

	products, err := s.repo.Get(ctx, dto.ProductSearchCriteria{CategoryID: categoryID})
	if err != nil {
		s.logger.Errorf("Error getting products: %v", err)
		return err
	}

	s.logger.Infof("Category: %v", category)
	s.logger.Infof("Products: %v", products)

	return nil
}
