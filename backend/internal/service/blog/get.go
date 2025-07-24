package blog

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
)

func (s *service) GetByID(ctx context.Context, id string) (*model.Blog, error) {
	blog, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if blog == nil {
		return nil, errors.New(404, "Blog not found")
	}

	return blog, nil
}

func (s *service) GetAll(ctx context.Context, offset, limit int) ([]model.Blog, error) {
	return s.repo.GetAll(ctx, offset, limit)
}
