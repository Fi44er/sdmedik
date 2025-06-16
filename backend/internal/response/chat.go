package response

import "github.com/Fi44er/sdmedik/backend/internal/model"

type ChatResponse struct {
	ID          string          `json:"id"`
	Messages    []model.Message `json:"messages"`
	UnreadCount int64           `json:"unread_count"`
}
