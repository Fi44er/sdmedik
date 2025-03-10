package basket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

func (s *service) AddItem(ctx context.Context, data *dto.AddBasketItem, userID string, sess *session.Session) error {
	if err := s.validator.Struct(data); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	var basket *model.Basket
	var err error

	if userID != "" {
		// Авторизованный пользователь → работаем с БД
		basket, err = s.repo.GetByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get basket by user ID: %w", err)
		}
		if basket == nil {
			return constants.ErrBasketNotFound
		}
	} else {
		// Гость → берем корзину из сессии
		if sess.Get("basket") == nil {
			basket = &model.Basket{}
		} else {
			str, ok := sess.Get("basket").(string)
			if !ok {
				return fmt.Errorf("Ошибка: interface{} не является строкой")
			}
			if err := json.Unmarshal([]byte(str), &basket); err != nil {
				return err
			}
		}
	}

	// Получаем товар
	product, _, err := s.productService.Get(ctx, dto.ProductSearchCriteria{ID: data.ProductID, Minimal: true})
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	if product == nil || len(*product) == 0 {
		return constants.ErrProductNotFound
	}

	// Проверяем, есть ли товар в корзине
	var basketItem *model.BasketItem

	if userID != "" {
		basketItem, err = s.basketItemRepo.GetByProductBasketID(ctx, data.ProductID, basket.ID)
		if err != nil {
			return fmt.Errorf("failed to get basket item: %w", err)
		}
	} else {
		// Гостевая корзина → ищем товар в сессионной корзине
		if basket != nil {
			for i, item := range basket.Items {
				if item.ProductID == data.ProductID {
					basketItem = &basket.Items[i]
					break
				}
			}
		}
	}

	if basketItem != nil {
		// Обновляем количество товара
		basketItem.Quantity += data.Quantity
		if basketItem.Quantity <= 0 {
			return s.DeleteItem(ctx, basketItem.ID, userID, sess) // Удаление товара
		}
		basketItem.TotalPrice = (*product)[0].Price * float64(basketItem.Quantity)

		if userID != "" {
			// Обновляем в БД
			if err := s.basketItemRepo.UpdateItemQuantity(ctx, basketItem); err != nil {
				return fmt.Errorf("failed to update basket item quantity: %w", err)
			}
		} else {
			s.logger.Infof("session is not nil: %+v", basket)
			// Обновляем в сессии
			for i, item := range basket.Items {
				if item.ProductID == basketItem.ProductID {
					basket.Items[i] = *basketItem
				}
			}

			basketStr, err := json.Marshal(basket)
			if err != nil {
				return fmt.Errorf("failed to marshal basket: %w", err)
			}
			sess.Set("basket", string(basketStr))
			sess.Save()
		}
		return nil
	}

	// Если товара нет в корзине, создаем новый элемент
	newBasketItem := model.BasketItem{
		Article:    (*product)[0].Article,
		Quantity:   data.Quantity,
		TotalPrice: (*product)[0].Price * float64(data.Quantity),
		ProductID:  data.ProductID,
	}

	if userID != "" {
		// Сохраняем в БД
		newBasketItem.BasketID = basket.ID
		if err := s.basketItemRepo.Create(ctx, &newBasketItem); err != nil {
			return fmt.Errorf("failed to create basket item: %w", err)
		}
	} else {
		// Добавляем в сессию
		newBasketItem.ID = uuid.NewString()
		basket.Items = append(basket.Items, newBasketItem)
		basketStr, err := json.Marshal(basket)
		if err != nil {
			return fmt.Errorf("failed to marshal basket: %w", err)
		}
		sess.Set("basket", string(basketStr))
		s.logger.Error("HUI")
		if err := sess.Save(); err != nil {
			return err
		}
	}

	return nil
}
