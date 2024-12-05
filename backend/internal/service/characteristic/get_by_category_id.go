package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByCategoryID(ctx context.Context, categoryID int) ([]model.Characteristic, error) {
	characteristics, err := s.repo.GetByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if len(characteristics) == 0 {
		return nil, errors.New(404, "characteristics not found")
	}

	return characteristics, nil
}
