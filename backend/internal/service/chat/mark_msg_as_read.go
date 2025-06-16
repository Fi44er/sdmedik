package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) MarkMsgAsRead(ctx context.Context, msgID string, userID string) error {
	message, err := s.repository.GetMessageByID(ctx, msgID)
	if err != nil {
		return err
	}

	if message == nil {
		return errors.New(404, "Message not found")
	}
	if message.SenderID != userID {
		return errors.New(403, "You are not allowed to mark this message as read")
	}

	return s.repository.MarkMsgAsRead(ctx, msgID)
}
