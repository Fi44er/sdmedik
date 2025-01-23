package webscraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model" // Предположим, что Certificate находится в этом пакете
	"github.com/Fi44er/sdmedik/backend/internal/service/webscraper/structs"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper"
)

func (s *service) Scraper() error {
	s.logger.Errorf("productService: %v", s.productService)
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

	s.logger.Logger.Infof("getManyCert: %v", getManyCert)

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
		options := utils.RequestOptions{
			Method: "GET",
			URL:    "https://esnsi.gosuslugi.ru/rest/ext/v1/classifiers/10616/data",
			Query: map[string]string{
				"query": item.CategoryName,
			},
		}
		esnsiRes, err := utils.MakeRequest(options)
		if err != nil {
			return err
		}

		var apiRes structs.ApiResponse
		if err := json.Unmarshal(esnsiRes, &apiRes); err != nil {
			fmt.Println("Ошибка при парсинге JSON:", err)
		}

		var tru string
		if len(apiRes.Body) > 0 {
			tru = apiRes.Body[0].Elements[3].Value
		}

		for _, region := range item.Items {
			key := fmt.Sprintf("%s-%s", item.CategoryArticle, region.Region)

			// Если запись существует, добавляем в updateCert
			if certMap[key] != "" {
				updateCert = append(updateCert, model.Certificate{
					ID:              certMap[key],
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
					Price:           region.Price,
					TRUName:         item.CategoryName,
					TRU:             tru,
				})
			} else {
				// Если записи нет, добавляем в createCert
				createCert = append(createCert, model.Certificate{
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
					Price:           region.Price,
					TRUName:         item.CategoryName,
					TRU:             tru,
				})
			}
		}

		for _, product := range item.Product {
			productDto := dto.CreateProduct{
				Article: product.Article,
				Name:    product.Name,
			}
			err := s.productService.Create(ctx, &productDto, nil)
			s.logger.Logger.Infof("productDto: %v", productDto)
			if err != nil {
				if !errors.Is(err, constants.ErrProductWithArticleConflict) {
					return err
				}
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
