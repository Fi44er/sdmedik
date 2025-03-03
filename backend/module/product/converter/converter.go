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

func ToDomainSlicceFromModel(models []model.Product) []domain.Product {
	domains := make([]domain.Product, len(models))
	for i, model := range models {
		domains[i] = *ToDomainFromModel(&model)
	}
	return domains
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
		CategoryIDs: dto.CategoryIDs,
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
	productsRes := make([]dto.ProductResponse, len(domains))
	for i, domain := range domains {
		productsRes[i] = *ToResponseFromDomain(&domain)
	}

	response := &dto.ProductsResponse{
		Products: productsRes,
		Count:    0,
	}
	return response
}

func ToModelProductCategoryFromDomain(domain *domain.ProductCategory) *model.ProductCategory {
	return &model.ProductCategory{
		CategoryID: domain.CategoryID,
		ProductID:  domain.ProductID,
	}
}

func ToModelProductCategorySliceFromDomain(domains []domain.ProductCategory) []model.ProductCategory {
	models := make([]model.ProductCategory, len(domains))
	for i, domain := range domains {
		models[i] = *ToModelProductCategoryFromDomain(&domain)
	}
	return models
}

func ToDomainProductCategoryFromModel(model *model.ProductCategory) *domain.ProductCategory {
	return &domain.ProductCategory{
		CategoryID: model.CategoryID,
		ProductID:  model.ProductID,
	}
}

func ToDomainProductCategorySliceFromModel(models []model.ProductCategory) []domain.ProductCategory {
	domains := make([]domain.ProductCategory, len(models))
	for i, model := range models {
		domains[i] = *ToDomainProductCategoryFromModel(&model)
	}
	return domains
}
