package webscraper

import (
	"context"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model" // Предположим, что Certificate находится в этом пакете
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper"
)

func (s *service) Scraper() error {
	s.logger.Errorf("CertService: %v", s.certificateService)
	ctx := context.Background()
	items := webscraper.Scraper()

	// Подготавливаем данные для поиска существующих сертификатов
	getManyCert := make([]dto.GetManyCert, 0)
	for _, item := range items {
		for _, region := range item.Items {
			getManyCert = append(getManyCert, dto.GetManyCert{
				CategoryArticle: item.CategoryArticle,
				RegionIso:       region.Region,
			})
		}
	}

	// Получаем существующие сертификаты из базы данных
	certs, err := s.certificateService.GetMany(ctx, &getManyCert)
	if err != nil {
		return err
	}

	// Создаём мапу для быстрого поиска существующих сертификатов
	certMap := make(map[string]string)
	for _, cert := range *certs {
		key := fmt.Sprintf("%s-%s", cert.CategoryArticle, cert.RegionIso)
		certMap[key] = cert.ID
	}

	// Разделяем данные на создание и обновление
	createCert := make([]model.Certificate, 0)
	updateCert := make([]model.Certificate, 0)

	for _, item := range items {
		for _, region := range item.Items {
			key := fmt.Sprintf("%s-%s", item.CategoryArticle, region.Region)

			// Если запись существует, добавляем в updateCert
			if certMap[key] != "" {
				updateCert = append(updateCert, model.Certificate{
					ID:              certMap[key],
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
					Price:           region.Price,
				})
			} else {
				// Если записи нет, добавляем в createCert
				createCert = append(createCert, model.Certificate{
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
					Price:           region.Price,
				})
			}
		}
	}

	// Создаём новые записи
	if len(createCert) > 0 {
		err := s.certificateService.CreateMany(ctx, &createCert)
		if err != nil {
			return fmt.Errorf("failed to create certificates: %v", err)
		}
	}

	// Обновляем существующие записи
	if len(updateCert) > 0 {
		err := s.certificateService.UpdateMany(ctx, &updateCert)
		if err != nil {
			return fmt.Errorf("failed to update certificates: %v", err)
		}
	}

	return nil
}
