package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/page"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	pageRepository "github.com/Fi44er/sdmedik/backend/internal/repository/page"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	pageService "github.com/Fi44er/sdmedik/backend/internal/service/page"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PageProvider struct {
	pageRepository repository.IPageRepository
	pageService    service.IPageService
	pageImpl       *page.Implementation

	db        *gorm.DB
	logger    *logger.Logger
	validator *validator.Validate
}

func NewPageProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
) *PageProvider {
	return &PageProvider{
		logger:    logger,
		db:        db,
		validator: validator,
	}
}

func (p *PageProvider) PageRepository() repository.IPageRepository {
	if p.pageRepository == nil {
		p.pageRepository = pageRepository.NewRepository(p.logger, p.db)
	}
	return p.pageRepository
}

func (p *PageProvider) PageService() service.IPageService {
	if p.pageService == nil {
		p.pageService = pageService.NewService(p.PageRepository(), p.logger, p.validator)
	}
	return p.pageService
}

func (p *PageProvider) PageImpl() *page.Implementation {
	if p.pageImpl == nil {
		p.pageImpl = page.NewImplementation(p.PageService())
	}
	return p.pageImpl
}
