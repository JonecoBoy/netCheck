package service

// Domain Service
// Encapsulates domain logic that operates on multiple entities.
// Application Service
// Coordinates tasks and orchestrates the use of domain services, entities, and repository.
type TaskService struct{}

func (s *TaskService) ScheduleTask(t *task.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) ReScheduleTask(t *task.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) deactivateTask(t *task.Task) error {
	// Domain logic for scheduling a task
	return nil
}

func (s *TaskService) cancelNextTaskOcurrency(t *task.Task) error {
	// Domain logic for scheduling a task
	return nil
}
