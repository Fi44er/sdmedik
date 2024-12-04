package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByID(ctx context.Context, id int) (model.Characteristic, error) {
	return s.repo.GetByID(ctx, id)
}
