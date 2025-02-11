package event

import (
	"time"
)

type SubscriptionEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

type SubscriptionCreated SubscriptionEvent

type SubscriptionUpdated SubscriptionEvent

type SubscriptionDeleted SubscriptionEvent

func (e *SubscriptionEvent) GetName() string {
	return e.name
}

func (e *SubscriptionEvent) GetDateTime() time.Time {
	return e.dateTime
}

func (e *SubscriptionEvent) GetPayload() interface{} {
	return e.payload
}
