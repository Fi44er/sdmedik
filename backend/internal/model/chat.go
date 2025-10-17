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
	ReadAt     time.Time `json:"read_at"`

	ChatID string `gorm:"type:string;not null" json:"chat_id"`
	Chat   Chat   `gorm:"foreignKey:ChatID" json:"-"`
	// Fragments []Fragment `gorm:"foreignKey:ChatID;references:ID" json:"fragments"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}

type Fragment struct {
	ID         string    `gorm:"primaryKey;type:string;" json:"id"`
	ChatID     string    `gorm:"type:string;not null;" json:"chat_id"`
	StartMsgID string    `gorm:"type:string;not null" json:"start_msg_id"`
	EndMsgID   *string   `gorm:"type:string" json:"end_msg_id"`
	Color      string    `gorm:"type:string;not null" json:"color"`
	CreatedAt  time.Time `gorm:"autoCreateTime;not null" json:"created_at"`

	// Add these relationships:
	StartMessage Message `gorm:"foreignKey:StartMsgID;references:ID"`
	EndMessage   Message `gorm:"foreignKey:EndMsgID;references:ID"`
	Chat         Chat    `gorm:"foreignKey:ChatID;references:ID"`
}

func (f *Fragment) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New().String()
	f.CreatedAt = time.Now()
	return nil
}

type Chat struct {
	ID        string     `gorm:"primaryKey;type:string;" json:"id"`
	Messages  []Message  `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;" json:"messages"`
	Fragments []Fragment `gorm:"foreignKey:ChatID;constraint:OnDelete:CASCADE;" json:"fragments"`
}
