package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func FindUsersByRole(repo repository.UserRepository, role int, page int, limit int) ([]*entity.User, error) {
	if err := entity.ValidateUserRole(role); err != nil {
		return nil, err
	}
	return repo.FindByRole(role, page, limit)
}
