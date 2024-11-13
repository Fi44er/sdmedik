package user

import "context"

func (s *service) Hello(ctx context.Context) string {
	s.logger.Info("Hello")
	return "Hello eee"
}
