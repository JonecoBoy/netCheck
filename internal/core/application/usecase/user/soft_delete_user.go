package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

func SoftDeleteUser(repo repository.UserRepository, id pkgEntity.ID) error {
	return repo.SoftDeleteByID(id)
}
