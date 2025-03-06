package basket

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (s *service) DeleteItem(ctx context.Context, itemID string, userID string, sess *session.Session) error {
	if userID == "" {
		basket, err := s.GetByUserID(ctx, userID, sess)
		if err != nil {
			return err
		}

		if err := s.basketItemRepo.Delete(ctx, itemID, basket.ID); err != nil {
			return err
		}
	} else {
		basket := sess.Get("basket").(*model.Basket)
		newBasket := []model.BasketItem{}
		for _, item := range basket.Items {
			if item.ID != itemID {
				newBasket = append(newBasket, item)
				break
			}
		}

		basket.Items = newBasket
		sess.Set("basket", basket)
		if err := sess.Save(); err != nil {
			return err
		}
	}

	return nil
}
