package chat

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IChatRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger, db *gorm.DB) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) GetAll(ctx context.Context, offset, limit int) ([]model.Chat, error) {
	r.logger.Info("Fetching chats...")

	// Нормализация параметров пагинации
	if offset <= 0 {
		offset = -1
	}
	if limit <= 0 {
		limit = -1 // или другое значение по умолчанию
	}

	var chats []model.Chat

	// Подзапрос для получения ID последних сообщений
	subQuery := r.db.Model(&model.Message{}).
		Select("chat_id, MAX(created_at) as max_created_at").
		Group("chat_id")

	// Основной запрос
	err := r.db.WithContext(ctx).
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN (?) as last_msg ON messages.chat_id = last_msg.chat_id AND messages.created_at = last_msg.max_created_at", subQuery)
		}).
		Offset(offset).
		Limit(limit).
		Order("chats.id").
		Find(&chats).Error

	if err != nil {
		r.logger.Errorf("Failed to fetch chats: %v", err)
		return nil, err
	}

	r.logger.Info("Chats fetched successfully")
	return chats, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Chat, error) {
	r.logger.Info("Fetching chat...")
	chat := new(model.Chat)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&chat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("Chat not found")
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch chat: %v", err)
		return nil, err
	}
	r.logger.Info("Chat fetched successfully")
	return chat, nil
}

func (r *repository) Create(ctx context.Context, chat *model.Chat) error {
	r.logger.Info("Creating chat...")
	if err := r.db.WithContext(ctx).Create(&chat).Error; err != nil {
		r.logger.Errorf("Failed to create chat: %v", err)
		return err
	}
	r.logger.Info("Chat created successfully")
	return nil
}

func (r *repository) SaveMessage(ctx context.Context, message *model.Message) error {
	r.logger.Info("Saving message...")
	r.logger.Infof("Message: %+v", message)
	if err := r.db.WithContext(ctx).Create(&message).Error; err != nil {
		r.logger.Errorf("Failed to save message: %v", err)
		return err
	}
	r.logger.Info("Message saved successfully")
	return nil
}

func (r *repository) GetMessagesByChatID(ctx context.Context, chatID string) ([]model.Message, error) {
	r.logger.Info("Fetching messages...")
	var messages []model.Message
	if err := r.db.WithContext(ctx).Where("chat_id = ?", chatID).Find(&messages).Error; err != nil {
		r.logger.Errorf("Failed to fetch messages: %v", err)
		return nil, err
	}
	r.logger.Info("Messages fetched successfully")
	return messages, nil
}

func (r *repository) GetMessageByID(ctx context.Context, id string) (*model.Message, error) {
	r.logger.Info("Fetching message...")
	message := new(model.Message)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&message).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Info("Message not found")
			return nil, nil
		}
		r.logger.Errorf("Failed to fetch message: %v", err)
		return nil, err
	}
	r.logger.Info("Message fetched successfully")
	return message, nil
}

func (r *repository) MarkMsgAsRead(ctx context.Context, msgID string) error {
	r.logger.Info("Marking message as read...")
	if err := r.db.WithContext(ctx).Model(&model.Message{}).Where("id = ?", msgID).Update("read_status", true).Error; err != nil {
		r.logger.Errorf("Failed to mark message as read: %v", err)
		return err
	}
	r.logger.Info("Message marked as read successfully")
	return nil
}

func (r *repository) GetUnreadCount(ctx context.Context, chatID, userID string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&model.Message{}).
		Where("chat_id = ? AND sender_id != ? AND read_status = ?", chatID, userID, false).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
