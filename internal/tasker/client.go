package tasker

import (
	"pierflow/internal/eventer"
)

type TaskClient interface {
	// Task executes a specific task defined in the Taskfile located in the project directory.
	//
	// on success, it sends a success message with the taskName, otherwise it sends an error message.
	Task(projectPath, taskFile, taskName string, messager eventer.Messager)

	// List lists all available tasks defined in the Taskfile located in the project directory.
	List(projectPath, taskFile string, messager eventer.Messager)
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
