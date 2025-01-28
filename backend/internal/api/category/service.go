package category

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	categoryService service.ICategoryService
}

func NewImplementation(categoryService service.ICategoryService) *Implementation {
	return &Implementation{
		categoryService: categoryService,
	}
}
