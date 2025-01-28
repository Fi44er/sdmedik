package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, dto *dto.CreateBasket) error {
	if err := s.validator.Struct(dto); err != nil {
		return err
	}

	basketModel := new(model.Basket)
	if err := utils.DtoToModel(dto, basketModel); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, basketModel); err != nil {
		return err
	}

	return nil
}
