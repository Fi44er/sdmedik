package product

import (
	"context"
	"strings"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
)

func (s *service) Get(ctx context.Context, criteria dto.ProductSearchCriteria) (*[]response.ProductResponse, *int64, error) {
	// Получаем продукты из репозитория
	products, count, err := s.repo.Get(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}

	// Создаем слайс для результата
	productRes := make([]response.ProductResponse, 0, len(*products))

	// Если требуется полная информация (не минимальная)
	var characteristicIDs []int
	var certificateDto []dto.GetManyCert
	for _, product := range *products {
		if criteria.Iso != "" {
			certificateDto = append(certificateDto, dto.GetManyCert{
				CategoryArticle: strings.Split(product.Article, ".")[0],
				RegionIso:       criteria.Iso,
			})
		}
		for _, characteristic := range product.CharacteristicValues {
			characteristicIDs = append(characteristicIDs, characteristic.CharacteristicID)
		}
	}

	// Получаем характеристики по их IDs
	characteristics, err := s.characteristicService.GetByIDs(ctx, characteristicIDs)
	if err != nil {
		return nil, nil, err
	}

	certificates, err := s.certificateService.GetMany(ctx, &certificateDto)
	if err != nil {
		return nil, nil, err
	}

	s.logger.Infof("certificates: %v", *certificates)

	certificateMap := make(map[string]float64)
	for _, certificate := range *certificates {
		certificateMap[certificate.CategoryArticle] = certificate.Price
	}

	// Создаем мапу характеристик для быстрого доступа по ID
	characteristicMap := make(map[int]model.Characteristic)
	for _, char := range *characteristics {
		characteristicMap[char.ID] = char
	}

	// Преобразуем каждый продукт в ProductResponse
	for _, product := range *products {
		// Преобразуем категории
		categoriesRes := make([]response.ProductCategoryRes, 0, len(product.Categories))
		for _, category := range product.Categories {
			categoriesRes = append(categoriesRes, response.ProductCategoryRes{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		// Преобразуем изображения
		imagesRes := make([]response.ProductImageRes, 0, len(product.Images))
		for _, image := range product.Images {
			imagesRes = append(imagesRes, response.ProductImageRes{
				ID:   image.ID,
				Name: image.Name,
			})
		}

		// Преобразуем характеристики
		characteristicsRes := make([]response.ProductCharacteristicRes, 0, len(product.CharacteristicValues))
		for _, charValue := range product.CharacteristicValues {
			if char, ok := characteristicMap[charValue.CharacteristicID]; ok {
				characteristicsRes = append(characteristicsRes, response.ProductCharacteristicRes{
					ID:     char.ID,
					Value:  charValue.Value,
					Name:   char.Name,
					Prices: charValue.Prices,
				})
			}
		}

		var certPrice float64
		if price, ok := certificateMap[strings.Split(product.Article, ".")[0]]; ok {
			certPrice = price
		} else {
			certPrice = 0
		}
		// Создаем ProductResponse
		productRes = append(productRes, response.ProductResponse{
			ID:               product.ID,
			Article:          product.Article,
			Name:             product.Name,
			Description:      product.Description,
			Price:            product.Price,
			Categories:       categoriesRes,
			Images:           imagesRes,
			Characteristic:   characteristicsRes,
			CertificatePrice: certPrice,
			Catalogs:         product.Catalogs,
			TRU:              product.TRU,
			Prewiew:          product.Preview,
			Nameplate:        product.Nameplate,
		})
	}

	return &productRes, count, nil
}
