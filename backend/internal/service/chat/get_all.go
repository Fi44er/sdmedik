package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetAll(ctx context.Context, offset, limit int) ([]model.Chat, error) {
	return s.repository.GetAll(ctx, offset, limit)
}
