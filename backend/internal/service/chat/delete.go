package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) DeleteChat(ctx context.Context, chatID string) error {
	s.logger.Warnf("Deleting chat with ID: %s", chatID)
	return s.repository.DeleteChat(ctx, chatID)
}

func (s *service) DeleteMessage(ctx context.Context, msgID, userID string) error {
	existMessage, err := s.repository.GetMessageByID(ctx, msgID)
	if err != nil {
		return err
	}

	if existMessage.SenderID != userID {
		return errors.New(403, "You are not delete this message")
	}

	return s.repository.DeleteMessageByID(ctx, msgID)
}

func (s *service) EditMessage(ctx context.Context, msg, msgID, userID string) error {
	existMessage, err := s.repository.GetMessageByID(ctx, msgID)
	if err != nil {
		return err
	}

	if existMessage.SenderID != userID {
		return errors.New(403, "You are not delete this message")
	}

	return s.repository.UpdateMessage(ctx, msgID, msg)
}
