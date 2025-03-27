package usecase

import (
	"time"

	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
)

func CreateTask(repo repository.TaskRepository, name, description string, taskType int, scheduledTime *time.Time, interval *int, crontab *string, command string) (*entity.Task, error) {
	task, err := entity.NewTask(name, description, taskType, scheduledTime, interval, crontab, command)
	if err != nil {
		return nil, err
	}
	err = repo.Save(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
