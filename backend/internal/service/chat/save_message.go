package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) SaveMessage(ctx context.Context, message *model.Message) error {
	if err := s.repository.SaveMessage(ctx, message); err != nil {
		return err
	}

	return nil
}
