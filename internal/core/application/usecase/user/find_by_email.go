package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func FindUserByEmail(repo repository.UserRepository, email string, page int, limit int) (*entity.User, error) {
	return repo.FindByEmail(email, page, limit)
}
