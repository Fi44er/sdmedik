package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) CreateMany(ctx context.Context, data *[]dto.CreateProduct) error {
	// Получаем все артикулы из входных данных
	articles := make([]string, len(*data))
	for i, product := range *data {
		articles[i] = product.Article
	}

	// Получаем существующие продукты из базы данных по артикулам
	existingProducts, err := s.repo.GetByArticles(ctx, articles)
	if err != nil {
		return err
	}

	// Создаем карту для существующих артикулов
	existingArticles := make(map[string]struct{})
	for _, product := range *existingProducts {
		existingArticles[product.Article] = struct{}{}
	}

	// Создаем новый срез для продуктов, которые будут добавлены
	var productsModels []model.Product

	for _, product := range *data {
		// Проверяем, существует ли продукт с таким артикулом
		if _, exists := existingArticles[product.Article]; exists {
			continue // Пропускаем, если артикул уже существует
		}

		productWithoutCharacteristic := product
		productWithoutCharacteristic.CharacteristicValues = nil
		var modelProduct model.Product
		if err := utils.DtoToModel(&productWithoutCharacteristic, &modelProduct); err != nil {
			return err
		}
		productsModels = append(productsModels, modelProduct)
	}

	// Создаем только те продукты, которые не существуют в базе данных
	if len(productsModels) > 0 {
		return s.repo.CreateMany(ctx, &productsModels)
	}

	return nil // Если нет новых продуктов для создания, возвращаем nil
}
