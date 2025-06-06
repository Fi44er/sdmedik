package certificate

import (
	"context"
	"errors"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	def "github.com/Fi44er/sdmedik/backend/internal/repository"
	"github.com/Fi44er/sdmedik/backend/pkg/logger"
	"gorm.io/gorm"
)

var _ def.ICertificateRepository = (*repository)(nil)

type repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger, db *gorm.DB) *repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) CreateMany(ctx context.Context, data *[]model.Certificate) error {
	r.logger.Info("Creating certificate...")
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Errorf("Failed to create certificate: %v", err)
		return err
	}
	r.logger.Infof("Certificate created successfully")
	return nil
}

func (r *repository) UpdateMany(ctx context.Context, data *[]model.Certificate) error {
	r.logger.Info("Updating certificates...")
	r.logger.Infof("Updating certificates: %v", data)

	for _, cert := range *data {
		if err := r.db.WithContext(ctx).Model(&model.Certificate{}).
			Where("id = ?", cert.ID).
			Updates(map[string]interface{}{
				"category_article": cert.CategoryArticle,
				"region_iso":       cert.RegionIso,
				"price":            cert.Price,
			}).Error; err != nil {
			r.logger.Errorf("Failed to update certificate: %v", err)
			return err
		}
	}

	r.logger.Info("Certificates updated successfully")
	return nil
}

func (r *repository) GetMany(ctx context.Context, data *[]dto.GetManyCert) (*[]model.Certificate, error) {
	r.logger.Info("Fetching certificates...")
	certificates := new([]model.Certificate)
	if len(*data) == 0 {
		return certificates, nil
	}

	// Сначала ищем точное совпадение по category_article + region_iso
	query := r.db.WithContext(ctx)
	for _, filter := range *data {
		query = query.Or("category_article = ? AND region_iso = ?", filter.CategoryArticle, filter.RegionIso)
	}

	if err := query.Find(&certificates).Error; err != nil {
		r.logger.Errorf("Failed to fetch certificates: %v", err)
		return nil, err
	}

	// Если ничего не найдено, ищем только по category_article (первое вхождение)
	if len(*certificates) == 0 {
		r.logger.Info("No exact matches found, trying with category_article only...")
		for _, filter := range *data {
			var fallbackCert model.Certificate
			err := r.db.WithContext(ctx).
				Where("category_article = ?", filter.CategoryArticle).
				First(&fallbackCert).Error

			if err == nil { // Найдена запись
				*certificates = append(*certificates, fallbackCert)
				break // Берём первую найденную и выходим
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				r.logger.Errorf("Error fetching fallback certificate: %v", err)
				return nil, err
			}
		}
	}

	r.logger.Info("Certificates fetched successfully")
	return certificates, nil
}
