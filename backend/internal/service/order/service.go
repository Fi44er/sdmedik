package order

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/Fi44er/sdmedik/backend/pkg/mailer"
	"github.com/go-playground/validator/v10"
)

var _ def.IOrderService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
	config    *config.Config
	repo      repository.IOrderRepository

	basketService  def.IBasketService
	certService    def.ICertificateService
	productService def.IProductService
	chatService    def.IChatService

	mailer *mailer.Mailer
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	config *config.Config,
	repo repository.IOrderRepository,
	basketService def.IBasketService,
	certService def.ICertificateService,
	productService def.IProductService,
	chatService def.IChatService,
) *service {
	templatePath := config.MailTemplatePath
	m, err := mailer.NewMailer(
		config.MailHost,
		config.MailPort,
		config.MailFrom,
		config.MailPassword,
		templatePath+"order.html",
		5,
	)

	if err != nil {
		logger.Fatalf("Failed to initialize mailer: %v", err)
	}

	return &service{
		logger:         logger,
		validator:      validator,
		config:         config,
		repo:           repo,
		basketService:  basketService,
		certService:    certService,
		productService: productService,
		mailer:         m,
		chatService:    chatService,
	}
}
