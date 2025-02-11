package event

import (
	"github.com/JonecoBoy/netCheck/internal/core/domain/event"
	"log"
)

type TaskEventPublisher struct{}

func (p *TaskEventPublisher) PublishTaskCreated(e event.TaskCreated) {
	log.Printf("Publishing TaskCreated event: %s", e.TaskID)
	// Logic to publish the event to a message broker
}

func (p *TaskEventPublisher) PublishTaskUpdated(e event.TaskUpdated) {
	log.Printf("Publishing TaskUpdated event: %s", e.TaskID)
	// Logic to publish the event to a message broker
}

func (p *TaskEventPublisher) PublishTaskDeleted(e event.TaskDeleted) {
	log.Printf("Publishing TaskDeleted event: %s", e.TaskID)
	// Logic to publish the event to a message broker
}
