package event

import (
	"time"
)

type DomainEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

type DomainCreated DomainEvent

type DomainUpdated DomainEvent

type DomainDeleted DomainEvent

func (e *DomainEvent) GetName() string {
	return e.name
}

func (e *DomainEvent) GetDateTime() time.Time {
	return e.dateTime
}

func (e *DomainEvent) GetPayload() interface{} {
	return e.payload
}
