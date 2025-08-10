package tasker

import (
	"context"
	"path"
	"pierflow/internal/logger"
	"time"

	"github.com/go-task/task/v3"
)

const RunTaskTimeout = 5 * time.Minute // RunTaskTimeout is the timeout for running a task (5 Minutes)

func (tc *taskClient) RunTask(ctx context.Context, projectPath, taskFile, taskName string) (string, error) {
	logger.Debugf("Taskfile: %s", taskFile)
	logger.Debugf("Project Path: %s", projectPath)
	logger.Debugf("Task Name: %s", taskName)

	readWriter := newTaskIO()
	executor := task.NewExecutor(
		task.WithEntrypoint(path.Join(tc.basePath, projectPath, taskFile)),
		task.WithDir(path.Join(tc.basePath, projectPath)),
		task.WithForce(true),
		task.WithVersionCheck(true),
		task.WithColor(false),            // Disable color output for better readability in logs
		task.WithIO(readWriter),          // Use the readWriter for input/output
		task.WithTimeout(RunTaskTimeout), // Set a timeout for the task execution
	)

	if err := executor.Setup(); err != nil {
		logger.Errorf("failed to setup executor: %s", err.Error())
		return "", err
	}

	err := executor.Run(ctx, &task.Call{Task: taskName, Silent: true})
	if err != nil {
		logger.Errorf("failed to run task: %s", err.Error())
		return "", err
	}
	return readWriter.String(), nil
}

func (tc *taskClient) TaskList(projectPath, taskFile string) ([]*TaskItem, error) {
	logger.Debugf("Taskfile: %s", taskFile)
	logger.Debugf("Project Path: %s", projectPath)

	readWriter := newTaskIO()
	executor := task.NewExecutor(
		task.WithEntrypoint(path.Join(tc.basePath, projectPath, taskFile)),
		task.WithDir(path.Join(tc.basePath, projectPath)),
		task.WithForce(true),
		task.WithVersionCheck(true),
		task.WithColor(false), // Disable color output for better readability in logs
		task.WithIO(readWriter),
	)

	if err := executor.Setup(); err != nil {
		logger.Errorf("failed to setup executor: %s", err.Error())
		return nil, err
	}

	tasks, err := executor.GetTaskList()
	if err != nil {
		logger.Errorf("failed to list tasks: %s", err.Error())
		return nil, err
	}
	var list []*TaskItem
	for _, taskItem := range tasks {
		list = append(list, toTaskItem(taskItem))
	}
	if len(list) == 0 {
		list = []*TaskItem{}
	}
	logger.Debugf("Task List output: %s", readWriter.String())
	return list, nil
}
