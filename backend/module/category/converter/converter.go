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

func ToDomainFromDTO(dto *dto.CreateCategoryDTO) *domain.Category {
	return &domain.Category{
		Name: dto.Name,
	}
}
