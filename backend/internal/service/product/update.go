package product

import (
	"context"
	"errors"
	"reflect"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
	custom_errors "github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Update(ctx context.Context, data *dto.UpdateProduct, images *dto.Images, id string) error {
	modelProduct := new(model.Product)
	var characteristicsValue []model.CharacteristicValue
	imagesID := []string{}
	imagesName := []string{}

	if err := s.validator.Struct(data); err != nil {
		return custom_errors.New(400, err.Error())
	}

	categories, err := s.categoryService.GetByIDs(ctx, data.CategoryIDs)
	if err != nil {
		return err
	}

	err = utils.ValidateCharacteristicValue(*categories, data.CharacteristicValues)
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

	modelProduct.ID = id

	dataValue := reflect.ValueOf(data).Elem()
	modelValue := reflect.ValueOf(modelProduct).Elem()
	for i := 0; i < dataValue.NumField(); i++ {
		field := dataValue.Field(i)
		fieldName := dataValue.Type().Field(i).Name

		if fieldName == "CharacteristicValues" {
			continue
		}

		if !field.IsZero() {
			modelField := modelValue.FieldByName(fieldName)
			if modelField.IsValid() && modelField.CanSet() {
				modelField.Set(field)
			}
		}
	}

	if err := s.repo.Update(ctx, modelProduct); err != nil {
		if errors.Is(err, constants.ErrProductNotFound) {
			s.transactionManagerRepo.Rollback(tx)
			return custom_errors.New(404, "Product not found")
		}
		return err
	}

	if err := s.characteristicValueService.DeleteByProductID(ctx, modelProduct.ID, tx); err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	for _, values := range data.CharacteristicValues {
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

	for _, image := range data.DelImages {
		imagesID = append(imagesID, image.ID)
		imagesName = append(imagesName, image.Name)
	}

	imageDto := dto.CreateImages{
		ProductID: modelProduct.ID,
		Images:    *images,
	}

	if len(images.Files) != 0 {
		if err := s.imageService.CreateMany(ctx, &imageDto, tx); err != nil {
			s.transactionManagerRepo.Rollback(tx)
			return err
		}
	}

	if err := s.imageService.DeleteByIDs(ctx, imagesID, imagesName, tx); err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	if err := s.transactionManagerRepo.Commit(tx); err != nil {
		return err
	}

	return nil
}
