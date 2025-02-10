package tasks

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task structure
type Task struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Type     string             `bson:"type"`     // "fixed" or "interval"
	Time     string             `bson:"time"`     // Only for "fixed" tasks
	Interval int                `bson:"interval"` // Only for "interval" tasks
	LastRun  time.Time          `bson:"last_run"`
	Command  string             `bson:"command"`
}
