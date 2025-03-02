package repository

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/module/category/converter"
	"github.com/Fi44er/sdmedik/backend/module/category/domain"
	"github.com/Fi44er/sdmedik/backend/module/category/model"
	"gorm.io/gorm"
)

func (r *CategoryRepository) GetByIDs(ctx context.Context, ids []string) ([]domain.Category, error) {
	r.logger.Infof("Getting categories by IDs: %v...", ids)
	var categories []model.Category
	err := r.db.Find(&categories, "id IN (?)", ids).Error
	if err != nil {
		r.logger.Errorf("Error getting categories by IDs: %v", err)
		return nil, err
	}
	categoryDomains := converter.ToDomainSliceFromModelSlice(categories)
	return categoryDomains, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	r.logger.Infof("Getting category by ID: %s...", id)
	var category model.Category
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Category with ID %s not found", id)
			return nil, nil
		}
		r.logger.Errorf("Error getting category by ID %v", err)
		return nil, err
	}

	files, err := r.fileRepo.GetFilesByOwner(ctx, category.ID, "category")
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		category.ImageIDs = append(category.ImageIDs, f.ID)
	}

	categoryDomain := converter.ToDomainFromModel(&category)

	return categoryDomain, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]domain.Category, error) {
	r.logger.Info("Getting all categories...")
	var categories []model.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		r.logger.Errorf("Error getting all categories: %v", err)
		return nil, err
	}

	categoryDomains := make([]domain.Category, len(categories))
	for i, category := range categories {
		files, err := r.fileRepo.GetFilesByOwner(ctx, category.ID, "category")
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			category.ImageIDs = append(category.ImageIDs, f.ID)
		}

		categoryDomains[i] = *converter.ToDomainFromModel(&category)
	}

	return categoryDomains, nil
}
