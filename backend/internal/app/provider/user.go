package provider

import (
	"github.com/Fi44er/sdmedik/backend/internal/api/user"
	"github.com/Fi44er/sdmedik/backend/internal/service"
	userService "github.com/Fi44er/sdmedik/backend/internal/service/user"
)

type UserProvider struct {
	// userRepository repository.UserRepository
	userService service.UserService
	userImpl    *user.Implementation
}

func newUserProvider() *UserProvider {
	return &UserProvider{}
}

// func (s *userProvider) UserRepository() repository.UserRepository {
// 	if s.userRepository == nil {
// 		s.userRepository = userRepository.NewRepository()
// 	}
// 	return s.userRepository
// }

func (s *UserProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService()
	}

	return s.userService
}

func (s *UserProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
	}

	return s.userImpl
}
