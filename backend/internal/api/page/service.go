package page

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	service service.IPageService
}

func NewImplementation(pageService service.IPageService) *Implementation {
	return &Implementation{
		service: pageService,
	}
}
