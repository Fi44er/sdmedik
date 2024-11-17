package product

import (
	"github.com/Fi44er/sdmedik/backend/internal/repository"
	def "github.com/Fi44er/sdmedik/backend/internal/service"
)

var _ def.IProductService = (*service)(nil)

type service struct {
	productRepository repository.IProductRepository
}

func NewService(productRepository repository.IProductRepository) *service {
	return &service{
		productRepository: productRepository,
	}
}
