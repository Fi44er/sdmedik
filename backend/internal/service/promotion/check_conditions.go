package promotion

import (
	"context"
	"strconv"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) CheckAndApplyPromotions(ctx context.Context, basket *response.BasketResponse) (*response.BasketResponse, error) {
	// Получаем все активные акции
	promotions, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Копируем корзину, чтобы не изменять оригинальную
	updatedBasket := &response.BasketResponse{
		ID:         basket.ID,
		Quantity:   basket.Quantity,
		TotalPrice: basket.TotalPrice,
		Items:      make([]response.BasketItemRes, len(basket.Items)),
	}
	copy(updatedBasket.Items, basket.Items)

	// Перебираем все акции

	for _, promotion := range *promotions {
		currentTime := time.Now().UTC().Add(5 * time.Hour)
		if promotion.EndDate.Before(currentTime) || promotion.StartDate.After(currentTime) {
			s.logger.Infof("currentTime: %v, startDate: %v, endDate: %v", currentTime, promotion.StartDate, promotion.EndDate)
			continue
		}

		// Проверяем условия акции
		conditionsMet := true

		switch promotion.Condition.Type {
		case model.ConditionTypeMinQuantity:
			// Проверка минимального количества товаров
			minQuantity, _ := strconv.Atoi(promotion.Condition.Value)
			totalQuantity := 0

			for _, item := range updatedBasket.Items {
				if item.ProductID == promotion.TargetID {
					totalQuantity += item.Quantity
				}
			}
			if totalQuantity < minQuantity {
				conditionsMet = false
				break
			}
		}

		// Если условия выполнены, применяем вознаграждение
		if conditionsMet {
			updatedBasket = s.applyReward(promotion, updatedBasket)
		}
	}

	return updatedBasket, nil
}

// ApplyReward применяет вознаграждение к корзине
func (s *service) applyReward(promotion model.Promotion, basket *response.BasketResponse) *response.BasketResponse {
	switch promotion.Reward.Type {
	case model.RewardTypePercentage:
		// Применение скидки в процентах
		for i := range basket.Items {
			if basket.Items[i].ProductID == promotion.TargetID {
				basket.Items[i].PriceWithPromotion = basket.Items[i].Price * (1 - promotion.Reward.Value/100)
				basket.Items[i].TotalPriceWithPromotion = basket.Items[i].PriceWithPromotion * float64(basket.Items[i].Quantity)
			}
		}

	case model.RewardTypeFixed:
		// Применение фиксированной скидки
		for i := range basket.Items {
			if basket.Items[i].ProductID == promotion.TargetID {
				basket.Items[i].PriceWithPromotion = basket.Items[i].Price - promotion.Reward.Value
				basket.Items[i].TotalPriceWithPromotion = basket.Items[i].PriceWithPromotion * float64(basket.Items[i].Quantity)
			}
		}

	case model.RewardTypeProduct:
		// Добавление бесплатного товара
		product, err := s.productService.GetByIDs(context.Background(), []string{promotion.GetProductID})
		if err != nil {
			s.logger.Errorf("Failed to get product: %v", err)
			return basket
		}

		var imageUrl string
		if len(*product) > 0 {
			if len((*product)[0].Images) > 0 {
				imageUrl = (*product)[0].Images[0].Name
			}
		}
		freeProduct := response.BasketItemRes{
			ID:         promotion.GetProductID, // ID товара, который даётся бесплатно
			ProductID:  promotion.GetProductID,
			Quantity:   int(promotion.Reward.Value), // Количество бесплатного товара
			Price:      0,                           // Бесплатно
			TotalPrice: 0,                           // Бесплатно
			IsFree:     true,
			Image:      imageUrl,
			Article:    (*product)[0].Article,
			Name:       (*product)[0].Name,
		}
		basket.Items = append(basket.Items, freeProduct)
	}

	// Пересчитываем общую стоимость корзины
	for _, item := range basket.Items {
		basket.TotalPriceWithPromotion += item.TotalPriceWithPromotion
	}

	return basket
}
