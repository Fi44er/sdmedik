package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) Create(ctx context.Context, chat *model.Chat) error {
	if chat.ID == "" {
		return errors.New(400, "Chat ID is required")
	}
	existChat, err := s.repository.GetByID(ctx, chat.ID)
	if err != nil {
		return err
	}

	if existChat == nil {
		return s.repository.Create(ctx, chat)
	}
	return nil
}
