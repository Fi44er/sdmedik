package webscraper

import (
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron"
)

var _ def.IWebScraperService = (*service)(nil)

type service struct {
	logger *logger.Logger
	// repo   repository.IWebScraperRepository
	validator *validator.Validate
	cron      *cron.Cron

	certificateService def.ICertificateService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	cron *cron.Cron,
	certificateService def.ICertificateService,
) *service {
	svc := &service{
		logger:             logger,
		validator:          validator,
		cron:               cron,
		certificateService: certificateService,
	}

	err := svc.cron.AddFunc("0 */5 * * * *", func() {
		svc.logger.Info("Задача выполняется каждые 5 минут")
		svc.Scraper()
	})
	if err != nil {
		svc.logger.Fatalf("Ошибка при добавлении задачи в cron: %v", err)
	}

	return svc
}
