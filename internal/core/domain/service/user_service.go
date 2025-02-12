package service

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
)

type UserService struct{}

func (s *UserService) CreateUser(user *entity.User) error {
	// Logic to create a user
	return nil
}
