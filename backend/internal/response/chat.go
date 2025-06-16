package response

import "github.com/Fi44er/sdmedik/backend/internal/model"

type ChatResponse struct {
	ID          string `gorm:"primaryKey;type:string;" json:"id"`
	Messages    []model.Message
	UnreadCount int64
}
