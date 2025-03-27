package repository

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"time"
)

type TaskRepository interface {
	FindFixedTasks(now time.Time) ([]*entity.Task, error)
	FindIntervalTasks(now time.Time) ([]*entity.Task, error)
	UpdateLastRun(taskID pkgEntity.ID) error
	Save(task *entity.Task) error
	DeleteByID(id pkgEntity.ID) error // ✅ Add this
	Update(task *entity.Task) error   // ✅ Add this
	SoftDeleteByID(id pkgEntity.ID) error
}
