package page

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) AddElement(ctx context.Context, data *dto.AddElement) error {
	page, err := s.repo.GetByPath(ctx, data.PagePath)
	if err != nil {
		return err
	}

	if page == nil {
		if err := s.repo.Create(ctx, &model.Page{
			Path: data.PagePath,
		}); err != nil {
			return err
		}
	}

	modelElement := model.Element{
		PagePath:  data.PagePath,
		ElementID: data.ElementID,
		Value:     data.Value,
	}

	if err := s.repo.AddOrUpdateElement(ctx, &modelElement); err != nil {
		return err
	}

	return nil
}
