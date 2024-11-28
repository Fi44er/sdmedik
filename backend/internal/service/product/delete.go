package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if err.Error() == "Product not found" {
			return errors.New(404, "Product not found")
		}
		return err
	}
	return nil
}
