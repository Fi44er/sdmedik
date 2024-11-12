package user

import def "github.com/Fi44er/sdmedik/backend/internal/service"

var _ def.UserService = (*service)(nil)

type service struct{}

func NewService() *service {
	return &service{}
}
