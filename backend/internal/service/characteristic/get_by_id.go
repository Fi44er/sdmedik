package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByID(ctx context.Context, id int) (model.Characteristic, error) {
	characteristic, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return model.Characteristic{}, err
	}

	if characteristic.ID == 0 {
		return model.Characteristic{}, errors.New(404, "characteristic not found")
	}

	return characteristic, nil
}