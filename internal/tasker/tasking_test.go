package tasker

import (
	"context"
	"pierflow/internal/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunTaskWith(t *testing.T) {
	_ = logger.InitLogLevel(logger.LogDebug)
	client := NewTaskClient(".")

	message, err := client.RunTask(context.Background(), "testdata", "Taskfile.yml", "test")
	assert.NoError(t, err, "Expected no error when running task")
	logger.Debug(message)
}

func TestTaskListWith(t *testing.T) {
	_ = logger.InitLogLevel(logger.LogDebug)
	client := NewTaskClient(".")

	tasks, err := client.TaskList("testdata", "Taskfile.yml")
	assert.NoError(t, err, "Expected no error when listing tasks")
	assert.NotEmpty(t, tasks, "Expected task list to not be empty")

	for _, task := range tasks {
		logger.Debugf("Task: %s - %s", task.Name, task.Desc)
	}
}
