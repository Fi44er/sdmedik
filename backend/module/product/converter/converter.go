package converter

import (
	"github.com/Fi44er/sdmedik/backend/module/product/domain"
	"github.com/Fi44er/sdmedik/backend/module/product/dto"
	"github.com/Fi44er/sdmedik/backend/module/product/model"
)

func ToDomainFromModel(model *model.Product) *domain.Product {
	return &domain.Product{
		ID:          model.ID,
		Article:     model.Article,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageIDs:    model.ImageIDs,
	}
}

func ToModelFromDomain(domain *domain.Product) *model.Product {
	return &model.Product{
		ID:          domain.ID,
		Article:     domain.Article,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		ImageIDs:    domain.ImageIDs,
	}
}

func ToDomainFromDTO(dto *dto.CreateProductDTO) *domain.Product {
	return &domain.Product{
		Article:     dto.Article,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       float64(dto.Price),
	}
}

func ToResponseFromDomain(domain *domain.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:          domain.ID,
		Article:     domain.Article,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		ImageIDs:    domain.ImageIDs,
	}
}

func ToResponseSliceFromDomain(domains []domain.Product) *dto.ProductsResponse {
	if domains == nil {
		return nil
	}

	productsRes := make([]dto.ProductResponse, len(domains))
	for i, d := range domains {
		productsRes[i] = *ToResponseFromDomain(&d)
	}

	response := &dto.ProductsResponse{
		Products: productsRes,
		Count:    0,
	}
	return response
}
