package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func CreateUser(repo repository.UserRepository, user *entity.User) error {
	return repo.Create(user)
}
