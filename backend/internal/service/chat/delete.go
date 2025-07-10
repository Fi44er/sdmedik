package chat

import "context"

func (s *service) DeleteChat(ctx context.Context, chatID string) error {
	s.logger.Warnf("Deleting chat with ID: %s", chatID)
	return s.repository.DeleteChat(ctx, chatID)
}
