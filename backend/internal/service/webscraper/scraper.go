package webscraper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model" // Предположим, что Certificate находится в этом пакете
	"github.com/Fi44er/sdmedik/backend/internal/service/webscraper/structs"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
	"github.com/Fi44er/sdmedik/backend/pkg/webscraper"
	scraper_structs "github.com/Fi44er/sdmedik/backend/pkg/webscraper/structs"
)

func (s *service) Scraper() error {
	s.mu.Lock()
	if s.cancelFunc != nil {
		s.cancelFunc() // Отменяем предыдущий процесс, если он выполнялся
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.cancelFunc = cancel
	s.mu.Unlock()

	s.logger.Errorf("productService: %v", s.productService)
	itemsChan := make(chan []scraper_structs.Items, 1)

	// Запускаем парсер в отдельной горутине
	go func() {
		items := webscraper.Scraper(ctx) // Новый парсер с контекстом
		itemsChan <- items
		close(itemsChan)
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("парсинг отменён")
	case items := <-itemsChan:
		getManyCert := make([]dto.GetManyCert, 0)
		for _, item := range items {
			for _, region := range item.Items {
				getManyCert = append(getManyCert, dto.GetManyCert{
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
				})
			}
		}

		chunks := s.chunkSlice(getManyCert, 1000)
		// Получаем существующие сертификаты из базы данных
		var allCerts []model.Certificate
		for _, chunk := range chunks {
			certs, err := s.certificateService.GetMany(ctx, &chunk)
			if err != nil {
				return err
			}
			allCerts = append(allCerts, *certs...)
		}

		// Создаём мапу для быстрого поиска существующих сертификатов
		certMap := make(map[string]string)
		for _, cert := range allCerts {
			key := fmt.Sprintf("%s-%s", cert.CategoryArticle, cert.RegionIso)
			certMap[key] = cert.ID
		}

		// Разделяем данные на создание и обновление
		createCert := make([]model.Certificate, 0)
		updateCert := make([]model.Certificate, 0)
		createProducts := make([]dto.CreateProduct, 0)

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
				createProducts = append(createProducts, productDto)
			}
		}

		productChank := s.chunkSliceProduct(createProducts, 1000)
		for _, chunk := range productChank {
			err := s.productService.CreateMany(ctx, &chunk)
			if err != nil {
				return err
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
}

func (s *service) chunkSlice(slice []dto.GetManyCert, chunkSize int) [][]dto.GetManyCert {
	var chunks [][]dto.GetManyCert
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func (s *service) chunkSliceProduct(slice []dto.CreateProduct, chunkSize int) [][]dto.CreateProduct {
	var chunks [][]dto.CreateProduct
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
