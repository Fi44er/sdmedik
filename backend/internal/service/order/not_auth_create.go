package order

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
)

func (s *service) NotAuthCreate(ctx context.Context, data *dto.CreateOrder, productID string) (string, error) {
	if err := s.validator.Struct(data); err != nil {
		return "", err
	}

	products, err := s.productService.GetByIDs(ctx, []string{productID})
	if err != nil {
		return "", err
	}
	s.logger.Infof("ID: %v, Products: %v", productID, *products)
	if len(*products) == 0 {
		return "", constants.ErrProductNotFound
	}

	product := (*products)[0]

	articles := make([]dto.GetManyCert, 1)
	categoryArticle := strings.Split(product.Article, ".")[0]
	articles = append(articles, dto.GetManyCert{CategoryArticle: categoryArticle})
	var image string
	if len(product.Images) > 0 {
		image = product.Images[0].Name
	}
	basket := response.BasketResponse{
		ID:         "1",
		Quantity:   1,
		TotalPrice: product.Price,
		Items: []response.BasketItemRes{
			{
				ID:         "1",
				Article:    product.Article,
				ProductID:  product.ID,
				Name:       product.Name,
				Image:      image,
				Quantity:   1,
				TotalPrice: product.Price,
				Price:      product.Price,
				IsFree:     false,
			},
		},
	}

	link, orderModel, err := s.sendToPaykeeper(ctx, data, &basket, articles, "")

	if err != nil {
		return "", err
	}

	orderItems := []model.OrderItem{
		{
			OrderID:    orderModel.ID,
			Name:       product.Name,
			Price:      product.Price,
			Quantity:   1,
			TotalPrice: product.Price,
			ProductID:  product.ID,
		},
	}

	if err := s.repo.AddItems(ctx, &orderItems); err != nil {
		return "", err
	}

	return link, nil
}
