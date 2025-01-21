package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/webscraper"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	webscraperService "github.com/Fi44er/sdmedik/backend/internal/service/webscraper"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron"
)

type WebscraperProvider struct {
	webScraperService service.IWebScraperService
	webScraperImpl    *webscraper.Implementation

	logger    *logger.Logger
	validator *validator.Validate
	cron      *cron.Cron

	certificateService service.ICertificateService
}

func NewWebscraperProvider(
	logger *logger.Logger,
	validator *validator.Validate,
	cron *cron.Cron,
	certificateService service.ICertificateService,
) *WebscraperProvider {
	return &WebscraperProvider{
		logger:             logger,
		validator:          validator,
		cron:               cron,
		certificateService: certificateService,
	}
}

func (p *WebscraperProvider) WebScraperService() service.IWebScraperService {
	if p.webScraperService == nil {
		p.webScraperService = webscraperService.NewService(p.logger, p.validator, p.cron, p.certificateService)
	}
	return p.webScraperService
}

func (p *WebscraperProvider) WebScraperImpl() *webscraper.Implementation {
	if p.webScraperImpl == nil {
		p.webScraperImpl = webscraper.NewImplementation(p.WebScraperService())
	}
	return p.webScraperImpl
}
