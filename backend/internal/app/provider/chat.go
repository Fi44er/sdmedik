package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/chat"
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"

	chatRepository "github.com/Fi44er/sdmedik/backend/internal/repository/chat"
	chatService "github.com/Fi44er/sdmedik/backend/internal/service/chat"
)

type ChatProvider struct {
	chatRepository repository.IChatRepository
	chatService    service.IChatService
	chatImpl       *chat.Implementation

	db     *gorm.DB
	logger *logger.Logger
	config *config.Config
}

func NewChatProvider(
	db *gorm.DB,
	logger *logger.Logger,
	config *config.Config,
) *ChatProvider {
	return &ChatProvider{
		db:     db,
		logger: logger,
		config: config,
	}
}

func (p *ChatProvider) ChatRepository() repository.IChatRepository {
	if p.chatRepository == nil {
		p.chatRepository = chatRepository.NewRepository(p.logger, p.db)
	}
	return p.chatRepository
}

func (p *ChatProvider) ChatService() service.IChatService {
	if p.chatService == nil {
		p.chatService = chatService.NewService(p.logger, p.ChatRepository(), p.config)
	}
	return p.chatService
}

func (p *ChatProvider) ChatImpl() *chat.Implementation {
	if p.chatImpl == nil {
		p.chatImpl = chat.NewImplementation(p.ChatService())
	}
	return p.chatImpl
}
