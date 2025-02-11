// application/usecase/create_task.go
package usecase

import (
	"github.com/JonecoBoy/netCheck/internal/core/service"
)

type CreateTaskUseCase struct {
	taskAppService *service.TaskAppService
}

func (uc *CreateTaskUseCase) Execute(t *task.Task) error {
	// Steps to create a task
	return uc.taskAppService.CreateTask(t)
}
