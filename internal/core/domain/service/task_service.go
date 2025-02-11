package service

import entity "github.com/jonecoboy/netCheck/internal/core/domain/entity/task"

// Domain Service
// Encapsulates domain logic that operates on multiple entities.
// Application Service
// Coordinates tasks and orchestrates the use of domain services, entities, and repository.
type TaskService struct{}

func (s *TaskService) ScheduleTask(t *entity.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) ReScheduleTask(t *entity.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) deactivateTask(t *entity.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) cancelNextTaskOcurrency(t *entity.Task) error {
	// Domain logic for scheduling a task
	return nil
}
