package user

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

func GetUser(repo repository.UserRepository, id pkgEntity.ID) (*entity.User, error) {
	return repo.FindByID(id)
}
