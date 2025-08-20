package tasker

import (
	"context"
)

type TaskClient interface {
	RunTask(ctx context.Context, projectPath, taskFile, taskName string, messageChan chan string)
	TaskList(projectPath, taskFile string) ([]*TaskItem, error)
}

type taskClient struct {
	basePath string
}

// NewTaskClient creates a new instance of TaskClient with the specified base path.
//
// The base path is used to determine the location of the project directory with the Taskfile.
// TaskFiles always reside in the project directory, so the base path should point to the root of the project.
func NewTaskClient(basePath string) TaskClient {
	return &taskClient{
		basePath: basePath,
	}
}
