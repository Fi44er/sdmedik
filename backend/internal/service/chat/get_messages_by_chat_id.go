package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetMessagesByChatID(ctx context.Context, chatID string) ([]model.Message, error) {
	return s.repository.GetMessagesByChatID(ctx, chatID)
}
