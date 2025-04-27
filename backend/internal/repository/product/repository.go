package product

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ def.IProductRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
	cache  *redis.Client
}

func NewRepository(
	logger *logger.Logger,
	db *gorm.DB,
	redis *redis.Client,
) *repository {
	return &repository{
		db:     db,
		logger: logger,
		cache:  redis,
	}
}

func (r *repository) DeleteCategoryAssociation(ctx context.Context, productID string, tx *gorm.DB) error {
	r.logger.Infof("Deleting category association for product with ID: %s...", productID)
	db := tx
	if db == nil {
		db = r.db
	}

	modelProduct := new(model.Product)
	if err := db.Preload("Categories").First(modelProduct, "id = ?", productID).Error; err != nil {
		r.logger.Errorf("Failed to fetch product: %v", err)
		return err
	}

	// Удаляем все текущие категории продукта
	if err := db.Model(modelProduct).Association("Categories").Clear(); err != nil {
		r.logger.Errorf("Failed to clear categories association: %v", err)
		return err
	}

	return nil
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

func (r *repository) Update(ctx context.Context, data *model.Product, tx *gorm.DB) error {
	r.logger.Info("Updating product...")

	db := tx
	if db == nil {
		db = r.db
	}
	if data.Catalogs == 0 {
		result := db.WithContext(ctx).Model(data).Update("catalogs", 0)
		if err := result.Error; err != nil {
			r.logger.Errorf("Failed to update product: %v", err)
			return err
		}
	}

	result := db.WithContext(ctx).Model(data).Updates(data)
	if err := result.Error; err != nil {
		r.logger.Errorf("Failed to update product: %v", err)
		return err
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

func (r *repository) Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]model.Product, *int64, error) {
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

	var count int64
	if criteria.CategoryID != 0 {
		err := r.db.Model(&model.Product{}).
			Joins("JOIN product_categories ON product_categories.product_id = products.id").
			Where("product_categories.category_id = ?", criteria.CategoryID).
			Count(&count).Error

		if err != nil {
			return nil, nil, err
		}
		request = request.Joins("JOIN product_categories ON product_categories.product_id = products.id").
			Where("product_categories.category_id = ?", criteria.CategoryID)
	} else {
		if err := r.db.Model(&model.Product{}).Count(&count).Error; err != nil {
			return nil, nil, err
		}
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
			// request = request.Joins(
			// 	"JOIN characteristic_values AS cv"+strconv.Itoa(filter.CharacteristicID)+" ON cv"+strconv.Itoa(filter.CharacteristicID)+".product_id = products.id",
			// ).Where(
			// 	"cv"+strconv.Itoa(filter.CharacteristicID)+".characteristic_id = ?", filter.CharacteristicID,
			// ).Where(
			// 	"cv"+strconv.Itoa(filter.CharacteristicID)+".value IN (?)", filter.Values,
			// )

			joinAlias := "cv" + strconv.Itoa(filter.CharacteristicID)

			request = request.Joins(
				"JOIN characteristic_values AS "+joinAlias+" ON "+joinAlias+".product_id = products.id",
			).Where(
				joinAlias+".characteristic_id = ?", filter.CharacteristicID,
			)

			if len(filter.Values) > 0 {
				jsonValue, _ := json.Marshal(filter.Values)

				// Явное приведение типа с обеих сторон оператора
				request = request.Where(
					joinAlias+".value::jsonb @> ?::jsonb", string(jsonValue),
				)
			}
		}
	}

	if criteria.Catalogs > 0 {
		var catalogMask uint8
		catalogMask = 1 << (criteria.Catalogs - 1)
		request = request.Where("catalogs & ? != 0", catalogMask)
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
			return products, nil, nil
		}
		r.logger.Errorf("Failed to fetch product: %v", err)
		return nil, nil, err
	}

	r.logger.Info("Product fetched successfully")
	return products, &count, nil
}

func (r *repository) GetByIDs(ctx context.Context, ids []string) (*[]model.Product, error) {
	products := new([]model.Product)
	if err := r.db.WithContext(ctx).Preload("Images").Where("id IN (?)", ids).Find(products).Error; err != nil {
		r.logger.Errorf("Failed to fetch products by IDs: %v", err)
		return nil, err
	}
	return products, nil
}

func (r *repository) GetTopProducts(ctx context.Context, limit int) ([]response.ProductPopularity, error) {
	var topProducts []response.ProductPopularity

	err := r.db.WithContext(ctx).
		Model(&model.OrderItem{}).
		Select("product_id, COUNT(*) as order_count").
		Group("product_id").
		Order("order_count DESC").
		Limit(limit).
		Find(&topProducts).Error

	if err != nil {
		return nil, err
	}

	return topProducts, nil
}

func (r *repository) CreateMany(ctx context.Context, data *[]model.Product) error {
	r.logger.Info("Creating products...")

	// Используем "ON CONFLICT DO NOTHING" для игнорирования дубликатов
	if err := r.db.WithContext(ctx).Table("products").Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create products: %v", err)
		return err
	}

	r.logger.Info("Products created successfully")
	return nil
}

func (r *repository) GetByArticles(ctx context.Context, articles []string) (*[]model.Product, error) {
	products := new([]model.Product)
	if err := r.db.WithContext(ctx).Where("article IN (?)", articles).Find(products).Error; err != nil {
		r.logger.Errorf("Failed to fetch products by articles: %v", err)
		return nil, err
	}
	return products, nil
}
