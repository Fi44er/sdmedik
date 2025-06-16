package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID         string    `gorm:"primaryKey;type:string;" json:"id"`
	SenderID   string    `gorm:"type:string;not null" json:"sender_id"`
	Message    string    `gorm:"type:text;not null" json:"message"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null" json:"time_to_send"`
	ReadStatus bool      `gorm:"not null;default:false" json:"read_status"`

	ChatID string `gorm:"type:string;not null" json:"chat_id"`
	Chat   Chat   `gorm:"foreignKey:ChatID" json:"-"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

type Chat struct {
	ID       string    `gorm:"primaryKey;type:string;" json:"id"`
	Messages []Message `gorm:"foreignKey:ChatID" json:"messages"`
}

// func (c *Chat) BeforeCreate(tx *gorm.DB) error {
// 	c.ID = uuid.New().String()
// 	return nil
// }
