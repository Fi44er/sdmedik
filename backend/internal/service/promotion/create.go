package promotion

import (
	"context"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
)

func (s *service) Create(ctx context.Context, data *dto.CreatePromotion) error {
	if err := s.validator.Struct(data); err != nil {
		return err
	}

	startDate, err := s.parseDate(data.StartDate)
	if err != nil {
		return err
	}
	s.logger.Info("111")

	endDate, err := s.parseDate(data.EndDate)
	if err != nil {
		s.logger.Errorf("Failed to parse end date: %v", err)
		return err
	}

	promotionModel := new(model.Promotion)
	if err := utils.DtoToModel(data, promotionModel); err != nil {
		s.logger.Errorf("Failed to convert dto to model: %v", err)
		return err
	}

	promotionModel.StartDate = startDate
	promotionModel.EndDate = endDate

	if err := s.repo.Create(ctx, promotionModel); err != nil {
		return err
	}

	return nil
}

func (s *service) parseDate(dateString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05" // Формат должен соответствовать строке

	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
