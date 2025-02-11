package event

import (
	"time"
)

type TaskEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

type TaskCreated TaskEvent

type TaskUpdated TaskEvent

type TaskDeleted TaskEvent

func (e *TaskEvent) GetName() string {
	return e.name
}

func (e *TaskEvent) GetDateTime() time.Time {
	return e.dateTime
}

func (e *TaskEvent) GetPayload() interface{} {
	return e.payload
}
