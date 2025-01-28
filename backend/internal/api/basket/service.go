package basket

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	basketService service.IBasketService
}

func NewImplementation(basketService service.IBasketService) *Implementation {
	return &Implementation{
		basketService: basketService,
	}
}
