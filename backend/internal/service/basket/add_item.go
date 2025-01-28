package basket

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) AddItem(ctx context.Context, data *dto.AddBasketItem, userID string) error {
	if err := s.validator.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	basket, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get basket by user ID: %w", err)
	}
	if basket == nil {
		return constants.ErrBasketNotFound
	}

	product, _, err := s.productService.Get(ctx, dto.ProductSearchCriteria{ID: data.ProductID, Minimal: true})
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	if product == nil || len(*product) == 0 {
		return constants.ErrProductNotFound
	}

	basketItem, err := s.basketItemRepo.GetByProductBasketID(ctx, data.ProductID, basket.ID)
	if err != nil {
		return fmt.Errorf("failed to get basket item: %w", err)
	}

	if basketItem != nil {
		s.logger.Infof("basketItem.Quantity: %v", basketItem.Quantity)

		basketItem.Quantity += data.Quantity
		if basketItem.Quantity <= 0 {
			if err := s.DeleteItem(ctx, basketItem.ID, userID); err != nil {
				return err
			}
			return nil
		}
		s.logger.Infof("basketItem.Quantity: %v", basketItem.Quantity)
		basketItem.TotalPrice = (*product)[0].Price * float64(basketItem.Quantity)
		if err := s.basketItemRepo.UpdateItemQuantity(ctx, basketItem); err != nil {
			return fmt.Errorf("failed to update basket item quantity: %w", err)
		}
		return nil
	} else {
		basketItemModel := &model.BasketItem{
			Article:    (*product)[0].Article,
			Quantity:   data.Quantity,
			TotalPrice: (*product)[0].Price * float64(data.Quantity),
			ProductID:  data.ProductID,
			BasketID:   basket.ID,
		}

		if err := s.basketItemRepo.Create(ctx, basketItemModel); err != nil {
			return fmt.Errorf("failed to create basket item: %w", err)
		}
	}

	return nil
}
