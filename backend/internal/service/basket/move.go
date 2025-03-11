package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
)

func (s *service) Move(ctx context.Context, data *dto.MoveBasket) error {
	_, err := s.GetByUserID(ctx, data.UserID, nil)
	if err != nil {
		return err
	}

	sessBasket, err := s.GetByUserID(ctx, "", data.Session)
	if err != nil {
		return err
	}

	for _, item := range sessBasket.Items {
		if err := s.AddItem(ctx, &dto.AddBasketItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}, data.UserID, nil); err != nil {
			return err
		}
	}

	if err := data.Session.Destroy(); err != nil {
		return err
	}

	return nil
}
