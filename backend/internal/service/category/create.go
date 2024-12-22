package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, data *dto.CreateCategory, image *dto.Image) error {
	if err := s.validator.Struct(data); err != nil {
		return errors.New(400, err.Error())
	}

	existProduct, err := s.repo.GetByName(ctx, data.Name)
	if err != nil {
		return err
	}

	if existProduct.Name != "" {
		return errors.New(409, "Category already exist")
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

	category := dto.CategoryWithoutCharacteristics{
		Name: data.Name,
	}

	var modelCategory model.Category
	if err := utils.DtoToModel(&category, &modelCategory); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, &modelCategory, tx); err != nil {
		s.logger.Errorf("Transaction rollback %v", err)
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	if len(data.Characteristics) != 0 {
		var characteristics []model.Characteristic
		for _, characteristic := range data.Characteristics {
			characteristics = append(characteristics, model.Characteristic{
				Name:       characteristic.Name,
				CategoryID: modelCategory.ID,
				DataType:   model.Type(characteristic.DataType),
			})
		}

		if err := s.characteristicService.CreateMany(ctx, &characteristics, tx); err != nil {
			s.transactionManagerRepo.Rollback(tx)
			s.logger.Errorf("Transaction rollback %v", err)
			return err
		}
	}

	images := new(dto.Images)
	images.Files = append(images.Files, image.File)
	imageDto := dto.CreateImages{
		CategoryID: modelCategory.ID,
		Images:     *images,
	}

	if err := s.imageService.CreateMany(ctx, &imageDto, tx); err != nil {
		s.transactionManagerRepo.Rollback(tx)
		return err
	}

	if err := s.transactionManagerRepo.Commit(tx); err != nil {
		return err
	}
	return nil
}
