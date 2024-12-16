package product

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"

	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, product *dto.CreateProduct) error {
	if err := s.validator.Struct(product); err != nil {
		return errors.New(400, err.Error())
	}

	existArticle, err := s.repo.Get(ctx, dto.ProductSearchCriteria{Article: product.Article})
	if err != nil {
		return err
	}

	if existArticle[0].ID != "" {
		return errors.New(400, "Product with this article already exists")
	}

	categories, err := s.categoryService.GetByIDs(ctx, product.CategoryIDs)
	if err != nil {
		return err
	}

	err = utils.ValidateCharacteristicValue(categories, product.CharacteristicValues)
	if err != nil {
		return err
	}

	tx, err := s.transactionManagerRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			s.logger.Error("Transaction rollback")
			s.transactionManagerRepo.Rollback(tx)
			panic(r) // Переподнимаем панику
		}
	}()

	productWithputCharacteristic := dto.Product{
		Article:     product.Article,
		Name:        product.Name,
		Description: product.Description,
	}

	var modelProduct model.Product
	if err := utils.DtoToModel(&productWithputCharacteristic, &modelProduct); err != nil {
		return err
	}

	modelProduct.Categories = categories

	if err := s.repo.Create(ctx, &modelProduct, tx); err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	var characteristicsValue []model.CharacteristicValue
	for _, values := range product.CharacteristicValues {
		characteristicsValue = append(characteristicsValue, model.CharacteristicValue{
			Value:            values.Value,
			CharacteristicID: values.CharacteristicID,
			ProductID:        modelProduct.ID,
		})
	}

	if len(characteristicsValue) != 0 {
		if err := s.characteristicValueService.CreateMany(ctx, &characteristicsValue, tx); err != nil {
			s.transactionManagerRepo.Rollback(tx)
			return err
		}
	}

	if err := s.transactionManagerRepo.Commit(tx); err != nil {
		return err
	}

	return nil
}
