package event

import (
	"time"
)

type UserEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

type UserCreated UserEvent

type UserUpdated UserEvent

type UserDeleted UserEvent

func (e *UserEvent) GetName() string {
	return e.name
}

func (e *UserEvent) GetDateTime() time.Time {
	return e.dateTime
}

func (e *UserEvent) GetPayload() interface{} {
	return e.payload
}
