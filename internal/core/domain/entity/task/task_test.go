package tasks

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
	task, err := newTask(name, description, Task_type_enum_fixed_time, &scheduledTime, nil, nil, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, Task_type_enum_fixed_time, task.Type)
	assert.Equal(t, scheduledTime, *task.ScheduledTime)
}
func TestNewTaskScheduledInterval(t *testing.T) {
	name := "Test Task"
	description := "This is a test task"
	command := "echo 'Hello, World!'"

	// Test Interval Task
	interval := 3600
	task, err := newTask(name, description, Task_type_enum_interval, nil, &interval, nil, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, Task_type_enum_interval, task.Type)
	assert.Equal(t, interval, *task.ScheduledInterval)

}
func TestNewTaskCronTab(t *testing.T) {
	name := "Test Task"
	description := "This is a test task"
	command := "echo 'Hello, World!'"

	// Test Crontab Task
	cronTab := "0 0 * * *"
	task, err := newTask(name, description, Task_type_enum_crontab, nil, nil, &cronTab, command)
	assert.Nil(t, err)
	assert.Equal(t, name, task.Name)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, command, task.Command)
	assert.Equal(t, Task_type_enum_crontab, task.Type)
	assert.Equal(t, cronTab, *task.CronTab)
}
