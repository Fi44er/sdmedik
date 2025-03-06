package basket

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (s *service) GetByUserID(ctx context.Context, userID string, sess *session.Session) (*response.BasketResponse, error) {
	basketRes := new(response.BasketResponse)
	var basket *model.Basket
	var err error
	if userID != "" {
		basket, err = s.repo.GetByUserID(ctx, userID)
		if err != nil {
			return nil, err
		}
	} else {
		if sess.Get("basket") != nil {
			str, ok := sess.Get("basket").(string)
			if !ok {
				return nil, fmt.Errorf("Ошибка: interface{} не является строкой")
			}
			if err := json.Unmarshal([]byte(str), &basket); err != nil {
				return nil, err
			}
		}
	}

	if basket == nil {
		return nil, constants.ErrBasketNotFound
	}

	var productsIDs []string
	for _, item := range basket.Items {
		productsIDs = append(productsIDs, item.ProductID)
	}

	products, err := s.productService.GetByIDs(ctx, productsIDs)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]model.Product)

	for _, product := range *products {
		productMap[product.ID] = product
	}

	totalPrice := 0.0
	totalQuantity := 0

	for _, item := range basket.Items {
		totalPrice += item.TotalPrice
		totalQuantity += item.Quantity
		var imageUrl string
		if len(productMap[item.ProductID].Images) > 0 {
			imageUrl = productMap[item.ProductID].Images[0].Name
		}
		basketRes.Items = append(basketRes.Items, response.BasketItemRes{
			ID:         item.ID,
			Article:    item.Article,
			ProductID:  item.ProductID,
			Name:       productMap[item.ProductID].Name,
			Image:      imageUrl,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			Price:      productMap[item.ProductID].Price,
		})
	}

	basketRes.ID = basket.ID
	basketRes.Quantity = totalQuantity
	basketRes.TotalPrice = math.Round(totalPrice*100) / 100
	promotionBasket, err := s.promotionService.CheckAndApplyPromotions(ctx, basketRes)
	if err != nil {
		return nil, err
	}

	return promotionBasket, nil
}
