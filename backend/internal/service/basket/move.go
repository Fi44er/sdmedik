package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
)

func (s *service) Move(ctx context.Context, data *dto.MoveBasket) error {
	s.logger.Infof("sess: %v", data.Session)

	newCtx := ctx
	_, err := s.GetByUserID(newCtx, data.UserID, nil)
	if err != nil {
		return err
	}

	sessBasket, err := s.GetByUserID(ctx, "", data.Session)
	if err != nil {
		return err
	}

	for _, item := range sessBasket.Items {
		newCtx := context.WithValue(ctx, "dinamic_options", item.SelectedOptions)
		if err := s.AddItem(newCtx, &dto.AddBasketItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Iso:       item.Iso,
		}, data.UserID, nil); err != nil {
			return err
		}
	}

	if err := data.Session.Destroy(); err != nil {
		return err
	}
	s.logger.Infof("sess2: %v", data.Session)

	return nil
}
