package category

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, data *dto.CreateCategory) error {
	s.logger.Info("Creating category in service...")

	if err := s.validator.Struct(data); err != nil {
		return errors.New(400, err.Error())
	}

	category := dto.CategoryWithoutCharacteristics{
		Name: data.Name,
	}

	var modelCategory model.Category
	if err := utils.DtoToModel(&category, &modelCategory); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, &modelCategory); err != nil {
		return err
	}

	if len(data.Characteristics) != 0 {
		var characteristics []model.Characteristic
		for _, characteristic := range data.Characteristics {
			characteristics = append(characteristics, model.Characteristic{
				Name:       characteristic,
				CategoryID: modelCategory.ID,
			})
		}

		if err := s.characteristicService.CreateMany(ctx, &characteristics); err != nil {
			return err
		}

	}
	s.logger.Info("Category created successfully")
	return nil
}
