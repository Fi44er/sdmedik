package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if err.Error() == "Category not found" {
			return errors.New(404, "Category not found")
		}
		return err
	}
	return nil
}
