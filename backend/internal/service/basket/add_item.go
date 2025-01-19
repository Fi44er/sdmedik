package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) AddItem(ctx context.Context, dto *dto.AddBasketItem) error {
	if err := s.validator.Struct(dto); err != nil {
		return err
	}

	basketItemModel := new(model.BasketItem)
	if err := utils.DtoToModel(dto, basketItemModel); err != nil {
		return err
	}
	if err := s.basketItemRepo.Create(ctx, basketItemModel); err != nil {
		return err
	}
	return nil
}
