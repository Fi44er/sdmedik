package basket

import "context"

func (s *service) DeleteItem(ctx context.Context, itemID string, userID string) error {
	basket, err := s.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if err := s.basketItemRepo.Delete(ctx, itemID, basket.ID); err != nil {
		return err
	}

	return nil
}
