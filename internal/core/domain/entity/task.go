package entity

import (
	"errors"
	"fmt"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/robfig/cron/v3"
	"time"
)

// Task structure
type Task struct {
	ID                pkgEntity.ID `json:"id" bson:"id"`
	User              User         `json:"user" bson:"user"`
	Name              string       `json:"name" bson:"name"`
	Description       string       `json:"description" bson:"description"`
	Type              int          `json:"type" bson:"type"`                             // "fixed" or "interval" or "crontab"
	ScheduledTime     *time.Time   `json:"scheduled_time" bson:"scheduled_time"`         // Only for "fixed" tasks
	ScheduledInterval *int         `json:"scheduled_interval" bson:"scheduled_interval"` // Only for "interval" tasks in seconds
	CronTab           *string      `json:"cron_tab" bson:"cron_tab"`                     // Only for "crontab" tasks
	LastRun           *time.Time   `json:"last_run" bson:"last_run"`
	Command           string       `json:"command" bson:"command"`
	CreatedAt         time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time    `bson:"updated_at" json:"updated_at"`
	DeletedAt         time.Time    `bson:"deleted_at" json:"deleted_at"`
}

const (
	_ = iota
	TaskTypeEnumFixedTime
	TaskTypeEnumInterval
	TaskTypeEnumCrontab
)

func NewTask(name string, description string, taskType int, time *time.Time, interval *int, crontab *string, command string) (*Task, error) {
	task := &Task{
		ID:                pkgEntity.NewId(),
		Name:              name,
		Description:       description,
		Type:              taskType,
		ScheduledTime:     time,
		ScheduledInterval: interval,
		CronTab:           crontab,
		Command:           command,
	}
	err := task.Validate()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func ValidateCronTab(cronTab string) error {
	_, err := cron.ParseStandard(cronTab)
	return err
}

func (t *Task) Validate() error {
	// validate task type
	switch t.Type {
	case TaskTypeEnumFixedTime:
		if t.ScheduledTime == nil {
			return errors.New("Fixed Time Task should have the ScheduledTime field filled")
		}
		if t.ScheduledInterval != nil && t.CronTab != nil {
			return errors.New("Fixed Time Task should have the fields ScheduledInterval and CronTab empty")
		}
		return nil
	case TaskTypeEnumInterval:
		if t.ScheduledInterval == nil {
			return errors.New("Interval Task should have the ScheduledInterval field filled")
		}
		if t.ScheduledTime != nil && t.CronTab != nil {
			return errors.New("Fixed Time Task should have the fields ScheduledTime and CronTab empty")
		}
		return nil
	case TaskTypeEnumCrontab:
		if t.CronTab == nil {
			return errors.New("Crontab Task should have the CronTab field filled")
		}
		if t.ScheduledTime != nil && t.ScheduledInterval != nil {
			return errors.New("Fixed Time Task should have the fields ScheduledTime and ScheduledInterval empty")
		}
		err := ValidateCronTab(*t.CronTab)
		if err != nil {
			return errors.New("Invalid CronTab string")
		}

		return nil
	default:
		return errors.New("invalid task type")
	}
}
func (t *Task) ShouldRun(now time.Time) bool {
	switch t.Type {
	case TaskTypeEnumFixedTime:
		return t.ScheduledTime != nil && now.Format("15:04") == t.ScheduledTime.Format("15:04")

	case TaskTypeEnumInterval:
		if t.ScheduledInterval == nil {
			return false
		}
		if t.LastRun == nil {
			return true
		}
		return time.Since(*t.LastRun) >= time.Duration(*t.ScheduledInterval)*time.Second

	case TaskTypeEnumCrontab:
		if t.CronTab == nil {
			return false
		}
		sched, err := cron.ParseStandard(*t.CronTab)
		if err != nil {
			return false
		}
		if t.LastRun == nil {
			return true
		}
		next := sched.Next(*t.LastRun)
		return now.After(next)
	}

	return false
}

func (t *Task) Execute() {
	// In real DDD, you'd dispatch the command to a handler
	fmt.Println("Executing:", t.Command)
}
