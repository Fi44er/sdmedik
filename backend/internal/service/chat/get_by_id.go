package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetByID(ctx context.Context, id string) (*model.Chat, error) {
	return s.repository.GetByID(ctx, id)
}
