package characteristic

import (
	"context"
	"errors"

	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	custom_errors "github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, constants.ErrCharacteristicNotFound) {
			return custom_errors.New(404, "Characteristic not found")
		}
		return err
	}

	return nil
}
