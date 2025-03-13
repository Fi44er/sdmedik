package webscraper

import (
	"context"
	"sync"

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
	productService     def.IProductService
	cancelFunc         context.CancelFunc
	mu                 sync.Mutex
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	cron *cron.Cron,
	certificateService def.ICertificateService,
	productService def.IProductService,
) *service {
	svc := &service{
		logger:             logger,
		validator:          validator,
		cron:               cron,
		certificateService: certificateService,
		productService:     productService,
	}

	// err := svc.cron.AddFunc("0 */3 * * * *", func() {
	// 	svc.logger.Info("Задача выполняется каждые 5 минут")
	// 	svc.Scraper()
	// })
	// if err != nil {
	// 	svc.logger.Fatalf("Ошибка при добавлении задачи в cron: %v", err)
	// }

	return svc
}
