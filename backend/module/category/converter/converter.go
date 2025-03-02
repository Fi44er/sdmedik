package converter

import (
	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	"github.com/Fi44er/sdmedik/backend/module/category/dto"
	"github.com/Fi44er/sdmedik/backend/module/category/model"
)

func ToModelFromDomain(domain *domain.Category) *model.Category {
	return &model.Category{
		ID:       domain.ID,
		Name:     domain.Name,
		ImageIDs: domain.ImageIDs,
	}
}

func ToDomainFromModel(model *model.Category) *domain.Category {
	return &domain.Category{
		ID:       model.ID,
		Name:     model.Name,
		ImageIDs: model.ImageIDs,
	}
}

func ToDomainSliceFromModelSlice(models []model.Category) []domain.Category {
	domains := make([]domain.Category, len(models))
	for i, model := range models {
		domains[i] = *ToDomainFromModel(&model)
	}
	return domains
}

func ToDomainFromDTO(dto *dto.CreateCategoryDTO) *domain.Category {
	return &domain.Category{
		Name: dto.Name,
	}
}

func ToResponseFromDomain(domain *domain.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:       domain.ID,
		Name:     domain.Name,
		ImageIDs: domain.ImageIDs,
	}
}

func ToResponseSliceFromDomainSlice(domains []domain.Category) []dto.CategoryResponse {
	response := make([]dto.CategoryResponse, len(domains))
	for i, domain := range domains {
		response[i] = *ToResponseFromDomain(&domain)
	}
	return response
}
