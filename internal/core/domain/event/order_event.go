package event

import (
	"time"
)

type OrderEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

type OrderCreated OrderEvent

type OrderUpdated OrderEvent

type OrderDeleted OrderEvent

func (e *OrderEvent) GetName() string {
	return e.name
}

func (e *OrderEvent) GetDateTime() time.Time {
	return e.dateTime
}

func (e *OrderEvent) GetPayload() interface{} {
	return e.payload
}
