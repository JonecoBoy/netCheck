package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTaskScheduledTime(t *testing.T) {
	name := "Test Task"
	description := "This is a test task"
	command := "echo 'Hello, World!'"

	// Test Fixed Time Task
	scheduledTime := time.Now().Add(1 * time.Hour)
	task, err := NewTask(name, description, TaskTypeEnumFixedTime, &scheduledTime, nil, nil, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, TaskTypeEnumFixedTime, task.Type)
	assert.Equal(t, scheduledTime, *task.ScheduledTime)
}
func TestNewTaskScheduledInterval(t *testing.T) {
	name := "Test Task"
	description := "This is a test task"
	command := "echo 'Hello, World!'"

	// Test Interval Task
	interval := 3600
	task, err := NewTask(name, description, TaskTypeEnumInterval, nil, &interval, nil, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, TaskTypeEnumInterval, task.Type)
	assert.Equal(t, interval, *task.ScheduledInterval)

}
func TestNewTaskCronTab(t *testing.T) {
	name := "Test Task"
	description := "This is a test task"
	command := "echo 'Hello, World!'"

	// Test Crontab Task
	cronTab := "0 0 * * *"
	task, err := NewTask(name, description, TaskTypeEnumCrontab, nil, nil, &cronTab, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, TaskTypeEnumCrontab, task.Type)
	assert.Equal(t, cronTab, *task.CronTab)
}
