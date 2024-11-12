package service

import "context"

type UserService interface {
	Hello(ctx context.Context) string
}
