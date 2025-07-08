package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) Move(ctx context.Context, data *dto.MoveBasket) error {
	_, err := s.GetByUserID(ctx, data.UserID, nil)
	if err != nil {
		return err
	}

	if data.Session == nil {
		return nil
	}

	sessBasket, err := s.GetByUserID(ctx, "", data.Session)
	if err != nil {
		if err != constants.ErrBasketNotFound {
			return err
		}
	}

	s.logger.Infof("sessBasket: %+v", sessBasket)

	if err == nil {
		for _, item := range sessBasket.Items {
			if err := s.AddItem(ctx, &dto.AddBasketItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Iso:       item.Iso,
			}, data.UserID, nil); err != nil {
				return err
			}
		}
	}

	return nil
}
