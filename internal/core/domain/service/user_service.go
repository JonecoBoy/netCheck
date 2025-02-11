package service

import entity "github.com/jonecoboy/netCheck/internal/core/domain/entity/user"

type UserService struct{}

func (s *UserService) CreateUser(user *entity.User) error {
	// Logic to create a user
	return nil
}
