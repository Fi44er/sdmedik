package webscraper

import "github.com/Fi44er/sdmedik/backend/internal/service"

type Implementation struct {
	webscraperService service.IWebScraperService
}

func NewImplementation(webscraperService service.IWebScraperService) *Implementation {
	return &Implementation{
		webscraperService: webscraperService,
	}
}
