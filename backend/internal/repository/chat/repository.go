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
	var chats []model.Chat
	if offset == 0 {
		offset = -1
	}

	if limit == 0 {
		limit = -1
	}

	if err := r.db.WithContext(ctx).
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("messages.created_at DESC").Limit(1)
		}).
		Offset(offset).
		Limit(limit).
		Find(&chats).Error; err != nil {
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
