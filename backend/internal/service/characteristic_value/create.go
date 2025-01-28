package characteristicvalue

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, characteristicValue *dto.CharacteristicValue) error {
	if err := s.validator.Struct(characteristicValue); err != nil {
		return errors.New(400, err.Error())
	}

	var modelCharacteristicValue model.CharacteristicValue
	if err := utils.DtoToModel(characteristicValue, &modelCharacteristicValue); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, &modelCharacteristicValue); err != nil {
		return err
	}

	return nil
}
