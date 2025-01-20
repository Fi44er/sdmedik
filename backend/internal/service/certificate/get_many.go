package certificate

import (
	"context"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
)

func (s *service) GetMany(ctx context.Context, data *[]dto.GetManyCert) (*[]model.Certificate, error) {
	s.logger.Info("Fetching certificates...")
	certificates, err := s.repo.GetMany(ctx, data)
	if err != nil {
		return nil, err
	}

	return certificates, nil
}
