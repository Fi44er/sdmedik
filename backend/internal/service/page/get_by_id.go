package page

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByPath(ctx context.Context, path string) (*model.Page, error) {
	page, err := s.repo.GetByPath(ctx, path)
	if err != nil {
		return nil, err
	}

	if page == nil {
		return nil, errors.New(404, "page not found")
	}
	return page, nil
}
