package chat

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
)

var _ def.IChatService = (*service)(nil)

type service struct {
	logger     *logger.Logger
	repository repository.IChatRepository
	mailer     *mailer.Mailer
	config     *config.Config
}

func NewService(
	logger *logger.Logger,
	repository repository.IChatRepository,
	config *config.Config,
) *service {
	temolatePath := config.MailTemplatePath
	m, err := mailer.NewMailer(
		config.MailHost,
		config.MailPort,
		config.MailFrom,
		config.MailPassword,
		temolatePath+"support.html",
		5,
	)

	if err != nil {
		logger.Fatalf("Failed to initialize mailer: %v", err)
	}

	return &service{
		logger:     logger,
		repository: repository,
		mailer:     m,
		config:     config,
	}
}
