package order

import (
	def "github.com/Fi44er/sdmedik/backend/internal/service"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/go-playground/validator/v10"
)

var _ def.IOrderService = (*service)(nil)

type service struct {
	logger    *logger.Logger
	validator *validator.Validate
}

func NewService(logger *logger.Logger, validator *validator.Validate) *service {
	return &service{
		logger:    logger,
		validator: validator,
	}
}
