package order

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Create(ctx context.Context, data *dto.CreateOrder, userID string) (string, error) {
	if err := s.validator.Struct(data); err != nil {
		return "", err
	}

	basket, err := s.basketService.GetByUserID(ctx, userID, nil)
	if err != nil {
		return "", err
	}

	articles := make([]dto.GetManyCert, len(basket.Items))
	for _, item := range basket.Items {
		categoryArticle := strings.Split(item.Article, ".")[0]
		articles = append(articles, dto.GetManyCert{CategoryArticle: categoryArticle})
	}

	link, orderModel, err := s.sendToPaykeeper(ctx, data, basket, articles, userID)
	if err != nil {
		return "", err
	}

	orderItems := []model.OrderItem{}
	for _, item := range basket.Items {
		orderItem := model.OrderItem{
			OrderID:    orderModel.ID,
			Name:       item.Name,
			Price:      item.Price,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice,
			ProductID:  item.ProductID,
		}
		orderItems = append(orderItems, orderItem)

		if err := s.basketService.DeleteItem(ctx, item.ID, userID, nil); err != nil {
			return "", err
		}
	}

	if err := s.repo.AddItems(ctx, &orderItems); err != nil {
		return "", err
	}

	return link, nil
}
