package product

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.IProductRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(
	logger *logger.Logger,
	db *gorm.DB,
) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, data *model.Product, tx *gorm.DB) error {
	r.logger.Info("Creating product...")

	db := tx
	if db == nil {
		db = r.db
	}

	if err := db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create product: %v", err)
		return err
	}

	r.logger.Infof("Product created successfully")
	return nil
}

func (r *repository) Update(ctx context.Context, data *model.Product) error {
	r.logger.Info("Updating product...")

	result := r.db.WithContext(ctx).Model(data).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update product: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Product with ID %s not found", data.ID)
		return fmt.Errorf("Product not found")
	}

	r.logger.Info("Product updated successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.logger.Infof("Deleting product with ID: %s...", id)
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete product: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Product with ID %s not found", id)
		return fmt.Errorf("Product not found")
	}

	r.logger.Infof("Product deleted by ID: %v successfully", id)
	return nil
}

func (r *repository) Get(ctx context.Context, criteria dto.ProductSearchCriteria) ([]model.Product, error) {
	var product []model.Product

	// Динамическое построение условий через рефлексию
	conditions := make(map[string]interface{})
	val := reflect.ValueOf(criteria)
	typ := reflect.TypeOf(criteria)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if !value.IsZero() { // Проверяем, заполнено ли поле
			conditions[field.Tag.Get("gorm")] = value.Interface()
		}
	}

	request := r.db.WithContext(ctx).Preload("Categories").Preload("Images")

	if criteria.CategoryID != 0 {
		request = request.Joins("JOIN product_categories ON product_categories.product_id = products.id").
			Where("product_categories.category_id = ?", criteria.CategoryID)
	}

	if criteria.Offset != 0 {
		request = request.Offset(criteria.Offset)
		delete(conditions, "offset")
	}

	if criteria.Limit != 0 {
		request = request.Limit(criteria.Limit)
		delete(conditions, "limit")
	}

	// Выполняем запрос с условиями
	r.logger.Infof("%v", conditions)
	err := request.Where(conditions).Find(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Product not found with provided criteria")
			return product, nil
		}
		r.logger.Errorf("Failed to fetch product: %v", err)
		return nil, err
	}

	r.logger.Info("Product fetched successfully")
	return product, nil
}
