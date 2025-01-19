package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) AddItem(ctx context.Context, data *dto.AddBasketItem, userID string) error {
	if err := s.validator.Struct(data); err != nil {
		return err
	}

	basket, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if basket == nil {
		return constants.ErrBasketNotFound
	}

	product, err := s.productService.Get(ctx, dto.ProductSearchCriteria{ID: data.ProductID, Minimal: true})
	if err != nil {
		return err
	}

	if len(*product) == 0 || product == nil {
		return constants.ErrProductNotFound
	}

	basketItem, err := s.basketItemRepo.GetByProductBasketID(ctx, data.ProductID, basket.ID)
	if err != nil {
		return err
	}
	if basketItem != nil {
		basketItem.Quantity = basketItem.Quantity + data.Quantity
		if err := s.basketItemRepo.UpdateItemQuantity(ctx, basketItem); err != nil {
			return err
		}

		return nil
	}

	basketItemModel := new(model.BasketItem)
	if err := utils.DtoToModel(data, basketItemModel); err != nil {
		return err
	}

	basketItemModel.BasketID = basket.ID

	if err := s.basketItemRepo.Create(ctx, basketItemModel); err != nil {
		return err
	}
	return nil
}
