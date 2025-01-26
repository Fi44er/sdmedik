package promotion

import (
	"context"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/constants"
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

	endDate, err := s.parseDate(data.EndDate)
	if err != nil {
		s.logger.Errorf("Failed to parse end date: %v", err)
		return err
	}

	if endDate.Before(startDate) {
		return constants.ErrPromotionDateBad
	}

	promotionWithoutRewardsConditions := dto.Promotion{
		Name:        data.Name,
		Description: data.Description,
		Type:        data.Type,
		TargetID:    data.TargetID,
	}

	promotionModel := new(model.Promotion)
	if err := utils.DtoToModel(&promotionWithoutRewardsConditions, promotionModel); err != nil {
		s.logger.Errorf("Failed to convert dto to model: %v", err)
		return err
	}

	promotionModel.StartDate = startDate
	promotionModel.EndDate = endDate

	if err := s.repo.Create(ctx, promotionModel); err != nil {
		return err
	}

	reward := model.Reward{
		PromotionID: promotionModel.ID,
		Type:        data.Reward.Type,
		Value:       data.Reward.Value,
	}
	if err := s.repo.CreateRewards(ctx, &reward); err != nil {
		return err
	}

	condition := model.Condition{
		PromotionID: promotionModel.ID,
		Type:        data.Condition.Type,
		Value:       data.Condition.Value,
	}
	if err := s.repo.CreateConditions(ctx, &condition); err != nil {
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
