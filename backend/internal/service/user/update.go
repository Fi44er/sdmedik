package user

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Update(ctx context.Context, data *dto.UpdateUser, id string) error {
	user := new(model.User)
	if err := utils.DtoToModel(data, user); err != nil {
		s.logger.Errorf("Error during conversion: %s", err.Error())
		return err
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
