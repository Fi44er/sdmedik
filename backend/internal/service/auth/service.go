package auth

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

var _ def.IAuthService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	cache     *redis.Client
	config    *config.Config
	mailer    *mailer.Mailer

	userService def.IUserService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	cache *redis.Client,
	userService def.IUserService,
) (*service, error) {
	m, err := mailer.NewMailer(
		config.MailHost,                  // SMTP-хост
		config.MailPort,                  // Порт
		config.MailFrom,                  // Ваш email
		config.MailPassword,              // Пароль от почты
		"pkg/mailer/template/index.html", // Путь к шаблону
		5,                                // Размер пула соединений
	)

	if err != nil {
		logger.Fatalf("Failed to initialize mailer: %v", err)
		return nil, err
	}

	return &service{
		logger:      logger,
		validator:   validator,
		config:      config,
		cache:       cache,
		mailer:      m,
		userService: userService,
	}, nil
}
