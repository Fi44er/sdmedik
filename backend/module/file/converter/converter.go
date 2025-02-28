package converter

import (
	"github.com/Fi44er/sdmedik/backend/module/file/domain"
	"github.com/Fi44er/sdmedik/backend/module/file/model"
)

func ToDomainFromModel(model *model.File) *domain.File {
	return &domain.File{
		ID:        model.ID,
		Name:      model.Name,
		OwnerID:   model.OwnerID,
		OwnerType: model.OwnerType,
	}
}

func ToModelFromDomain(domain *domain.File) *model.File {
	return &model.File{
		ID:        domain.ID,
		Name:      domain.Name,
		OwnerID:   domain.OwnerID,
		OwnerType: domain.OwnerType,
	}
}

func ToDomainSliceFromModel(models []model.File) []domain.File {
	if models == nil {
		return nil
	}

	domains := make([]domain.File, len(models))
	for i, m := range models {
		domains[i] = *ToDomainFromModel(&m)
	}
	return domains
}

func ToModelSliceFromDomain(domains []domain.File) []model.File {
	if domains == nil {
		return nil
	}

	models := make([]model.File, len(domains))
	for i, d := range domains {
		models[i] = *ToModelFromDomain(&d)
	}
	return models
}
