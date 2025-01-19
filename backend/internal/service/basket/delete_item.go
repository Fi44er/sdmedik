package basket

import "context"

func (s *service) DeleteItem(ctx context.Context, itemID string) error {
	if err := s.validator.Struct(itemID); err != nil {
		return err
	}
	return nil
}
