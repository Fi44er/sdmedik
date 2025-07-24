package blog

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	service service.IBlogService
}

func NewImplementation(service service.IBlogService) *Implementation {
	return &Implementation{service: service}
}
