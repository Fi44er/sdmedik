package certificate

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

var _ def.ICertificateService = (*service)(nil)

type service struct {
	repo   repository.ICertificateRepository
	logger *logger.Logger
}

func NewService(repo repository.ICertificateRepository, logger *logger.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
