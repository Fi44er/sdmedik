package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByIDs(ctx context.Context, ids []int) (*[]model.Characteristic, error) {
	characteristics, err := s.repo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return characteristics, nil
}
