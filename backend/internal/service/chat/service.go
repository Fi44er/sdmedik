package chat

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

var _ def.IChatService = (*service)(nil)

type service struct {
	logger     *logger.Logger
	repository repository.IChatRepository
}

func NewService(
	logger *logger.Logger,
	repository repository.IChatRepository,
) *service {
	return &service{
		logger:     logger,
		repository: repository,
	}
}
