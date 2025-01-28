package search

import (
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.ISearchService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate

	indexService def.IIndexService
}

func NewService(
	logger *logger.Logger,
	validator *validator.Validate,
	indexService def.IIndexService,
) *service {
	return &service{
		logger:       logger,
		validator:    validator,
		indexService: indexService,
	}
}
