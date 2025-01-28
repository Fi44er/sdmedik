package order

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	orderService service.IOrderService
}

func NewImplementation(orderService service.IOrderService) *Implementation {
	return &Implementation{
		orderService: orderService,
	}
}
