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
		basket, err = s.repo.GetByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get basket by user ID: %w", err)
		}
		if basket == nil {
			return constants.ErrBasketNotFound
		}
	} else {
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

	product, _, err := s.productService.Get(ctx, dto.ProductSearchCriteria{ID: data.ProductID, Minimal: true, Iso: data.Iso})
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	if product == nil || len(*product) == 0 {
		return constants.ErrProductNotFound
	}

	var catalogMask uint8 = 1 << 1
	isSertificate := (*product)[0].Catalogs&catalogMask != 0 && data.Iso != ""

	var basketItem *model.BasketItem
	if userID != "" {
		// Ищем только тот товар, который совпадает по `ProductID` и `IsSertificate`
		basketItem, err = s.basketItemRepo.GetByProductIDIsoIsCert(ctx, data.ProductID, basket.ID, data.Iso, isSertificate)
		if err != nil {
			return fmt.Errorf("failed to get basket item: %w", err)
		}
	} else {
		for i, item := range basket.Items {
			if item.ProductID == data.ProductID && item.IsCertificate == isSertificate && item.Iso == data.Iso {
				basketItem = &basket.Items[i]
				break
			}
		}
	}

	if basketItem != nil {
		// Обновляем количество товара
		basketItem.Quantity += data.Quantity
		if basketItem.Quantity <= 0 {
			return s.DeleteItem(ctx, basketItem.ID, userID, sess)
		}

		if isSertificate {
			basketItem.TotalPrice = (*product)[0].CertificatePrice * float64(basketItem.Quantity)
		} else {
			basketItem.TotalPrice = (*product)[0].Price * float64(basketItem.Quantity)
		}

		if userID != "" {
			if err := s.basketItemRepo.UpdateItemQuantity(ctx, basketItem); err != nil {
				return fmt.Errorf("failed to update basket item quantity: %w", err)
			}
		} else {
			s.logger.Infof("session is not nil: %+v", basket)
			for i, item := range basket.Items {
				if item.ProductID == basketItem.ProductID && item.IsCertificate == isSertificate {
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

	// Создаем новый элемент в корзине
	var totalPrice float64
	iso := ""
	if isSertificate {
		totalPrice = (*product)[0].CertificatePrice * float64(data.Quantity)
		iso = data.Iso
	} else {
		totalPrice = (*product)[0].Price * float64(data.Quantity)
	}

	newBasketItem := model.BasketItem{
		Article:       (*product)[0].Article,
		Quantity:      data.Quantity,
		TotalPrice:    totalPrice,
		ProductID:     data.ProductID,
		IsCertificate: isSertificate,
		Iso:           iso,
	}

	if userID != "" {
		newBasketItem.BasketID = basket.ID
		if err := s.basketItemRepo.Create(ctx, &newBasketItem); err != nil {
			return fmt.Errorf("failed to create basket item: %w", err)
		}
	} else {
		newBasketItem.ID = uuid.NewString()
		basket.Items = append(basket.Items, newBasketItem)
		basketStr, err := json.Marshal(basket)
		if err != nil {
			return fmt.Errorf("failed to marshal basket: %w", err)
		}
		sess.Set("basket", string(basketStr))
		if err := sess.Save(); err != nil {
			return err
		}
	}

	return nil
}
