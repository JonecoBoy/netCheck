package usecase

import (
	"time"

	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func UpdateTask(repo repository.TaskRepository, task *entity.Task) error {
	task.UpdatedAt = time.Now()

	if err := task.Validate(); err != nil {
		return err
	}

	return repo.Update(task)
}
