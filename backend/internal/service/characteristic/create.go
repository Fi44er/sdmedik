package characteristic

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, data *dto.CreateCharacteristic) error {
	if err := s.validator.Struct(data); err != nil {
		return errors.New(400, err.Error())
	}

	modelCharacteristic := new(model.Characteristic)
	if err := utils.DtoToModel(data, modelCharacteristic); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, modelCharacteristic); err != nil {
		return err
	}

	return nil
}
