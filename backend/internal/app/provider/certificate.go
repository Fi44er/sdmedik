package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	certificateRepository "github.com/Fi44er/sdmedik/backend/internal/repository/certificate"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	certificateService "github.com/Fi44er/sdmedik/backend/internal/service/certificate"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

type CertificateProvider struct {
	certificateRepository repository.ICertificateRepository
	certificateService    service.ICertificateService

	logger *logger.Logger
	db     *gorm.DB
}

func NewCertificateProvider(
	logger *logger.Logger,
	db *gorm.DB,
) *CertificateProvider {
	return &CertificateProvider{
		logger: logger,
		db:     db,
	}
}

func (p *CertificateProvider) CertificateRepository() repository.ICertificateRepository {
	if p.certificateRepository == nil {
		p.certificateRepository = certificateRepository.NewRepository(p.logger, p.db)
	}
	return p.certificateRepository
}

func (p *CertificateProvider) CertificateService() service.ICertificateService {
	if p.certificateService == nil {
		p.certificateService = certificateService.NewService(p.CertificateRepository(), p.logger)
	}
	return p.certificateService
}
