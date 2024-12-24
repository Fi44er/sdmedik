package characteristicvalue

import (
	"context"

	"gorm.io/gorm"
)

func (s *service) DeleteByProductID(ctx context.Context, productID string, tx *gorm.DB) error {
	if err := s.repo.DeleteByProductID(ctx, productID, tx); err != nil {
		return err
	}
	return nil
}
