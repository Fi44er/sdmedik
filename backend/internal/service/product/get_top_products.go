package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) GetTopProducts(ctx context.Context, limit int) (*[]response.TopProductRes, error) {
	top, err := s.repo.GetTopProducts(ctx, limit)
	if err != nil {
		return nil, err
	}

	productsIds := []string{}

	for _, item := range top {
		productsIds = append(productsIds, item.ProductID)
	}

	products, err := s.repo.GetByIDs(ctx, productsIds)
	if err != nil {
		return nil, err
	}

	topProductList := []response.TopProductRes{}
	for index, product := range *products {
		var imageUrl string
		if len(product.Images) > 0 {
			imageUrl = product.Images[0].Name
		}
		topProductList = append(topProductList, response.TopProductRes{
			ID:         product.ID,
			Price:      product.Price,
			OrderCount: top[index].OrderCount,
			Article:    product.Article,
			Image:      imageUrl,
			Name:       product.Name,
		})
	}

	return &topProductList, nil
}
