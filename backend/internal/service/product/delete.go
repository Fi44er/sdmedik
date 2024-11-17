package product

import "context"

func (s *service) Delete(ctx context.Context, id string) error {
	return s.productRepository.Delete(ctx, id)
}
