package user

import (
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

var _ def.IUserService = (*service)(nil)

type service struct {
	logger *logger.Logger
}

func NewService(logger *logger.Logger) *service {
	return &service{
		logger: logger,
	}
}
