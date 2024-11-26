package product

import "context"

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.productRepository.Delete(ctx, id); err != nil {
		s.logger.Errorf("Failed to delete product: %v", err)
		return err
	}
	return nil
}
