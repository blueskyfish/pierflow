package tasker

import (
	"context"
	"path"
	"pierflow/internal/logger"
	"time"

	"github.com/go-task/task/v3"
)

const RunTaskTimeout = 5 * time.Minute // RunTaskTimeout is the timeout for running a task (5 Minutes)

func (tc *taskClient) RunTask(ctx context.Context, projectPath, taskFile, taskName string, messageChan chan string) {
	logger.Debugf("Taskfile: %s", taskFile)
	logger.Debugf("Project Path: %s", projectPath)
	logger.Debugf("Task Name: %s", taskName)

	go tc.runTask(ctx, projectPath, taskFile, taskName, messageChan)
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

func (tc *taskClient) runTask(ctx context.Context, projectPath, taskFile, taskName string, messageChan chan string) {

	defer close(messageChan)

	outWriter := newTaskWriter("okay", messageChan)
	executor := task.NewExecutor(
		task.WithEntrypoint(path.Join(tc.basePath, projectPath, taskFile)),
		task.WithDir(path.Join(tc.basePath, projectPath)),
		task.WithForce(true),
		task.WithVersionCheck(true),
		task.WithColor(false),            // Disable color output for better readability in logs
		task.WithStdout(outWriter),       // Use the provided writer for stdout
		task.WithStderr(outWriter),       // Use the provided writer for stderr
		task.WithTimeout(RunTaskTimeout), // Set a timeout for the task execution
	)

	if err := executor.Setup(); err != nil {
		logger.Errorf("failed to setup executor: %s", err.Error())
		messageChan <- toErrorf("failed to setup executor: %s", err.Error())
		return
	}

	err := executor.Run(ctx, &task.Call{Task: taskName, Silent: true})
	if err != nil {
		logger.Errorf("failed to run task: %s", err.Error())
		messageChan <- toErrorf("failed to run task: %s", err.Error())
	} else {
		messageChan <- toOkayf("Task %s finished", taskName)
	}
}
