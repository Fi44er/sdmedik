package search

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	searchService service.ISearchService
}

func NewImplementation(searchService service.ISearchService) *Implementation {
	return &Implementation{
		searchService: searchService,
	}
}
