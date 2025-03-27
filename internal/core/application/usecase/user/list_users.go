package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func ListUsers(repo repository.UserRepository, page int, limit int) ([]*entity.User, error) {
	return repo.FindAll(page, limit)
}
