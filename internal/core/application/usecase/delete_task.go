package usecase

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

func DeleteTask(repo repository.TaskRepository, id pkgEntity.ID) error {
	return repo.DeleteByID(id)
}
