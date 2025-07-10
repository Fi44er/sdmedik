package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) GetMessagesByChatID(ctx context.Context, chatID string) ([]response.FragmenResponse, error) {
	fragments, err := s.repository.GetFragmentsByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	chatRes := make([]response.FragmenResponse, 0)
	for _, fragment := range fragments {
		messages, err := s.repository.GetMessagesInFragment(ctx, fragment)
		if err != nil {
			return nil, err
		}

		chatRes = append(chatRes, response.FragmenResponse{
			ID:       fragment.ID,
			Color:    fragment.Color,
			Messages: messages,
		})
	}

	return chatRes, nil
}
