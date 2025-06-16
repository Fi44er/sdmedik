package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) GetAll(ctx context.Context, offset, limit int, userID string) ([]response.ChatResponse, error) {
	chats, err := s.repository.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var res []response.ChatResponse
	for chatsIndex := range chats {
		chat := chats[chatsIndex]
		unreadCount, err := s.repository.GetUnreadCount(ctx, chat.ID, userID)
		if err != nil {
			return nil, err
		}
		res = append(res, response.ChatResponse{
			ID:          chat.ID,
			Messages:    chat.Messages,
			UnreadCount: unreadCount,
		})
	}

	return res, nil
}
