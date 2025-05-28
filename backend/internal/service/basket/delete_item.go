package basket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (s *service) DeleteItem(ctx context.Context, itemID string, userID string, sess *session.Session) error {
	if userID != "" {
		basket, err := s.GetByUserID(ctx, userID, sess)
		if err != nil {
			return err
		}

		if err := s.basketItemRepo.Delete(ctx, itemID, basket.ID); err != nil {
			return err
		}
	} else {
		var basket model.Basket
		str, ok := sess.Get("basket").(string)
		if !ok {
			return fmt.Errorf("Ошибка: interface{} не является строкой")
		}
		if err := json.Unmarshal([]byte(str), &basket); err != nil {
			return err
		}

		newBasket := []model.BasketItem{}
		for _, item := range basket.Items {
			if item.ID != itemID {
				newBasket = append(newBasket, item)
				break
			}
		}

		basket.Items = newBasket

		basketStr, err := json.Marshal(basket)
		if err != nil {
			return fmt.Errorf("failed to marshal basket: %w", err)
		}

		sess.Set("basket", string(basketStr))
	}

	return nil
}
