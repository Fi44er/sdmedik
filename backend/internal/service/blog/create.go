package blog

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) Create(ctx context.Context, data *model.Blog) error {
	return s.repo.Create(ctx, data)
}

func (s *service) Update(ctx context.Context, data *model.Blog) error {
	return s.repo.Update(ctx, data)
}
