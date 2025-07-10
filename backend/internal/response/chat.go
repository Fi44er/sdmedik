package response

import "github.com/Fi44er/sdmedik/backend/internal/model"

type ChatResponse struct {
	ID          string          `json:"id"`
	Messages    []model.Message `json:"messages"`
	UnreadCount int64           `json:"unread_count"`
}

type FragmenResponse struct {
	ID       string          `json:"id"`
	Color    string          `json:"color"`
	Messages []model.Message `json:"messages"`
}
