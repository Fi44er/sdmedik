package order

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) ChangeStatus(ctx context.Context, data *dto.ChangeOrderStatus) error {
	if err := s.validator.Struct(data); err != nil {
		return err
	}

	orderModel := model.Order{
		ID:     data.OrderID,
		Status: data.Status,
	}

	if err := s.repo.Update(ctx, &orderModel); err != nil {
		return err
	}

	return nil
}
