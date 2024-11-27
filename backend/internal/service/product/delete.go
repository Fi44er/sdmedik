package product

import "context"

func (s *service) Delete(ctx context.Context, id string) error {
	s.logger.Info("Deleting product in service...")
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Errorf("Failed to delete product: %v", err)
		return err
	}
	return nil
}
