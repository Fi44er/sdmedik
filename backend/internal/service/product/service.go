package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
)

var _ def.IProductService = (*service)(nil)

type service struct {
	logger            *logger.Logger
	productRepository repository.IProductRepository
}

func NewService(productRepository repository.IProductRepository, logger *logger.Logger) *service {
	return &service{
		productRepository: productRepository,
		logger:            logger,
	}
}
