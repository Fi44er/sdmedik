package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/config"
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	imageRepository "github.com/Fi44er/sdmedik/backend/internal/repository/image"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	imageService "github.com/Fi44er/sdmedik/backend/internal/service/image"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ImageProvider struct {
	imageRepository repository.IImageRepository
	imageServise    service.IImageService

	logger    *logger.Logger
	db        *gorm.DB
	validator *validator.Validate
	config    *config.Config
}

func NewImageProvider(
	logger *logger.Logger,
	db *gorm.DB,
	validator *validator.Validate,
	config *config.Config,
) *ImageProvider {
	return &ImageProvider{
		logger:    logger,
		db:        db,
		validator: validator,
		config:    config,
	}
}

func (p *ImageProvider) ImageRepository() repository.IImageRepository {
	if p.imageRepository == nil {
		p.imageRepository = imageRepository.NewRepository(p.logger, p.db)
	}

	return p.imageRepository
}

func (p *ImageProvider) ImageService() service.IImageService {
	if p.imageServise == nil {
		p.imageServise = imageService.NewService(p.logger, p.validator, p.ImageRepository(), p.config)
	}
	return p.imageServise
}
