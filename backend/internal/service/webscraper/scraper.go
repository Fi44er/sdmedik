package webscraper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
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

	s.logger.Info("Starting web scraping process...")

	itemsChan := make(chan []scraper_structs.Items, 1)

	// Запускаем парсер в отдельной горутине
	go func() {
		items := webscraper.Scraper(ctx) // Новый парсер с контекстом
		itemsChan <- items
		close(itemsChan)
	}()

	select {
	case <-ctx.Done():
		s.logger.Warn("Парсинг отменён")
		return fmt.Errorf("парсинг отменён")
	case items := <-itemsChan:
		s.logger.Infof("Получено %d элементов для обработки", len(items))
		getManyCert := make([]dto.GetManyCert, 0)

		for _, item := range items {
			s.logger.Infof("Обработка элемента: %s", item.CategoryArticle)
			for _, region := range item.Items {
				getManyCert = append(getManyCert, dto.GetManyCert{
					CategoryArticle: item.CategoryArticle,
					RegionIso:       region.Region,
				})
			}
		}

		chunks := s.chunkSlice(getManyCert, 1000)
		s.logger.Infof("Получаем существующие сертификаты из базы данных, всего пакетов: %d", len(chunks))
		var allCerts []model.Certificate
		for _, chunk := range chunks {
			certs, err := s.certificateService.GetMany(ctx, &chunk)
			if err != nil {
				s.logger.Errorf("Ошибка при получении сертификатов: %v", err)
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
			s.logger.Infof("Запрос к API для категории: %s", item.CategoryName)
			options := utils.RequestOptions{
				Method: "GET",
				URL:    "https://esnsi.gosuslugi.ru/rest/ext/v1/classifiers/10616/data",
				Query: map[string]string{
					"query": item.CategoryName,
				},
			}
			esnsiRes, err := utils.MakeRequest(options)
			if err != nil {
				s.logger.Errorf("Ошибка при выполнении запроса: %v", err)
				return err
			}

			var apiRes structs.ApiResponse
			if err := json.Unmarshal(esnsiRes, &apiRes); err != nil {
				s.logger.Errorf("Ошибка при парсинге JSON: %v", err)
				return err
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
					s.logger.Infof("Обновление сертификата: %s для региона: %s", certMap[key], region.Region)
				} else {
					// Если записи нет, добавляем в createCert
					createCert = append(createCert, model.Certificate{
						CategoryArticle: item.CategoryArticle,
						RegionIso:       region.Region,
						Price:           region.Price,
						TRUName:         item.CategoryName,
						TRU:             tru,
					})
					s.logger.Infof("Создание нового сертификата для категории: %s, регион: %s", item.CategoryArticle, region.Region)
				}
			}

			for _, product := range item.Product {
				productDto := dto.CreateProduct{
					Article: product.Article,
					Name:    product.Name,
				}
				createProducts = append(createProducts, productDto)
				s.logger.Infof("Добавление продукта: %s", product.Name)
			}
		}

		// Создаём новые записи
		createCertChunk := s.chunckSliceCert(updateCert, 1000)
		for _, chunk := range createCertChunk {
			err := s.certificateService.CreateMany(ctx, &chunk)
			if err != nil {
				s.logger.Errorf("Ошибка при обновлении сертификатов: %v", err)
				return fmt.Errorf("failed to update certificates: %v", err)
			}
			s.logger.Infof("Успешно обновлено %d сертификатов", len(chunk))
		}

		// Обновляем существующие записи
		updateCertChunk := s.chunckSliceCert(updateCert, 1000)
		for _, chunk := range updateCertChunk {
			err := s.certificateService.UpdateMany(ctx, &chunk)
			if err != nil {
				s.logger.Errorf("Ошибка при обновлении сертификатов: %v", err)
				return fmt.Errorf("failed to update certificates: %v", err)
			}
			s.logger.Infof("Успешно обновлено %d сертификатов", len(chunk))
		}

		productChunk := s.chunkSliceProduct(createProducts, 1000)
		for _, chunk := range productChunk {
			err := s.productService.CreateMany(ctx, &chunk)
			if err != nil {
				s.logger.Errorf("Ошибка при создании продуктов: %v", err)
				return err
			}
			s.logger.Infof("Успешно создано %d продуктов", len(chunk))
		}

		s.logger.Info("Парсинг и обработка завершены успешно")
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

func (s *service) chunckSliceCert(slice []model.Certificate, chunkSize int) [][]model.Certificate {
	var chunks [][]model.Certificate
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
