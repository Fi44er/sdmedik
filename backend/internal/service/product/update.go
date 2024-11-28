package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Update(ctx context.Context, product *model.Product) error {

	if err := s.repo.Update(ctx, product); err != nil {
		if err.Error() == "Product not found" {
			return errors.New(404, "Product not found")
		}
		return err
	}

	return nil
}
