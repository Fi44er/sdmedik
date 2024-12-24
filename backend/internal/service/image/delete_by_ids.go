package image

import (
	"context"

	"gorm.io/gorm"
)

func (s *service) DeleteByIDs(ctx context.Context, ids []string, names []string, tx *gorm.DB) error {
	if err := s.repo.DeleteByIDs(ctx, ids, tx); err != nil {
		return err
	}

	if err := s.DeleteByNames(ctx, names); err != nil {
		s.logger.Errorf("Error deleting files: %v", err)
		return err
	}

	return nil
}
