package product

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	productService service.IProductService
}

func NewImplementation(productService service.IProductService) *Implementation {
	return &Implementation{
		productService: productService,
	}
}
