package usecase

import (
	"github.com/jonecoboy/netCheck/internal/core/domain/repository"
	"log"
	"time"
)

func StartTaskProcessor(repo repository.TaskRepository) {
	go func() {
		for {
			tasks, err := repo.FindFixedTasks(time.Now())
			if err != nil {
				log.Println("Fixed task error:", err)
				continue
			}
			for _, task := range tasks {
				if task.ShouldRun(time.Now()) {
					task.Execute()
					repo.UpdateLastRun(task.ID)
				}
			}
			time.Sleep(60 * time.Second)
		}
	}()

	go func() {
		for {
			tasks, err := repo.FindIntervalTasks(time.Now())
			if err != nil {
				log.Println("Interval task error:", err)
				continue
			}
			for _, task := range tasks {
				if task.ShouldRun(time.Now()) {
					task.Execute()
					repo.UpdateLastRun(task.ID)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
