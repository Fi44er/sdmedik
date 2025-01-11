package product

import (
	"context"
	"reflect"
	"strconv"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
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
		return constants.ErrProductNotFound
	}

	r.logger.Info("Product updated successfully")
	return nil
}

func (r *repository) Delete(ctx context.Context, id string, tx *gorm.DB) error {
	r.logger.Infof("Deleting product with ID: %s...", id)
	db := tx
	if db == nil {
		r.logger.Error("Transaction is nil")
		db = r.db
	}
	result := db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{})
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to delete product: %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		r.logger.Warnf("Product with ID %s not found", id)
		return constants.ErrProductNotFound
	}

	r.logger.Infof("Product deleted by ID: %v successfully", id)
	return nil
}

func (r *repository) Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]model.Product, error) {
	products := new([]model.Product)

	conditions := make(map[string]interface{})
	val := reflect.ValueOf(criteria)
	typ := reflect.TypeOf(criteria)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if field.Tag.Get("gorm") == "-" {
			continue
		}

		if !value.IsZero() { // Проверяем, заполнено ли поле
			conditions[field.Tag.Get("gorm")] = value.Interface()
		}
	}

	request := r.db.WithContext(ctx)
	if !criteria.Minimal {
		request = request.Preload("Categories").Preload("Images").Preload("CharacteristicValues")
	}

	if criteria.CategoryID != 0 {
		request = request.Joins("JOIN product_categories ON product_categories.product_id = products.id").
			Where("product_categories.category_id = ?", criteria.CategoryID)
	}

	if criteria.Filters.Price.Min > 0 || criteria.Filters.Price.Max > 0 {
		if criteria.Filters.Price.Min > 0 {
			request = request.Where("price >= ?", criteria.Filters.Price.Min)
		}
		if criteria.Filters.Price.Max > 0 {
			request = request.Where("price <= ?", criteria.Filters.Price.Max)
		}
	}

	if len(criteria.Filters.Characteristics) > 0 {
		for _, filter := range criteria.Filters.Characteristics {
			request = request.Joins(
				"JOIN characteristic_values AS cv"+strconv.Itoa(filter.CharacteristicID)+" ON cv"+strconv.Itoa(filter.CharacteristicID)+".product_id = products.id",
			).Where(
				"cv"+strconv.Itoa(filter.CharacteristicID)+".characteristic_id = ?", filter.CharacteristicID,
			).Where(
				"cv"+strconv.Itoa(filter.CharacteristicID)+".value IN (?)", filter.Values,
			)
		}
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
	err := request.Where(conditions).Find(products).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warnf("Product not found with provided criteria")
			return products, nil
		}
		r.logger.Errorf("Failed to fetch product: %v", err)
		return nil, err
	}

	r.logger.Info("Product fetched successfully")
	return products, nil
}
