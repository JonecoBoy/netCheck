//go:generate mockgen -source=user_repository.go -destination=../repository/mocks/user_repository_mock.go -package=mocks

package repository

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	SoftDeleteByID(id pkgEntity.ID) error
	DeleteByID(id pkgEntity.ID) error
	FindByID(id pkgEntity.ID) (*entity.User, error)
	Find(filter string) ([]*entity.User, error)
	FindAll(page, limit int) ([]*entity.User, error)
	FindByRole(role int, page int, limit int) ([]*entity.User, error)
	FindByEmail(email string, page int, limit int) (*entity.User, error)
}
