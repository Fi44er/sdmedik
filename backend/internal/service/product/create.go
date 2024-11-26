package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Create(ctx context.Context, product *model.Product) error {
	s.logger.Info("Creating product in service...")

	if err := s.productRepository.Create(ctx, product); err != nil {
		s.logger.Errorf("Failed to create product: %v", err)
		return err
	}

	s.logger.Info("Product created successfully")
	return nil
}
