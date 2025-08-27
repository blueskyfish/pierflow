package tasker

import (
	"context"
	"fmt"
	"io"
	"path"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"

	"github.com/go-task/task/v3"
)

func (tc *taskClient) RunTask(projectPath, taskFile, taskName string, messager eventer.Messager) {
	logger.Debugf("Taskfile: %s", taskFile)
	logger.Debugf("Project Path: %s", projectPath)
	logger.Debugf("Task Name: %s", taskName)

	go tc.runTask(projectPath, taskFile, taskName, messager)
}

func (tc *taskClient) TaskList(projectPath, taskFile string, messager eventer.Messager) {
	logger.Debugf("Taskfile: %s", taskFile)
	logger.Debugf("Project Path: %s", projectPath)
	go tc.runTaskList(projectPath, taskFile, messager)
}

func (tc *taskClient) runTask(projectPath, taskFile, taskName string, messager eventer.Messager) {
	defer messager.Closing()

	executor := tc.createExecute(projectPath, taskFile, messager)
	if err := executor.Setup(); err != nil {
		logger.Errorf("failed to setup executor: %s", err.Error())
		_ = messager.Send(eventer.StatusError, fmt.Sprintf("failed to setup executor: %s", err.Error()))
		return
	}

	ctx := context.Background()
	err := executor.Run(ctx, &task.Call{Task: taskName, Silent: true})
	if err != nil {
		logger.Errorf("failed to run task: %s", err.Error())
		_ = messager.Send(eventer.StatusError, fmt.Sprintf("failed to run task: %s", err.Error()))
	} else {
		_ = messager.Send(eventer.StatusInfo, fmt.Sprintf("task [%s] finished", taskName))
	}
}

func (tc *taskClient) runTaskList(projectPath, taskFile string, messager eventer.Messager) {
	defer messager.Closing()

	executor := tc.createExecute(projectPath, taskFile, messager)
	if err := executor.Setup(); err != nil {
		logger.Errorf("failed to setup executor: %s", err.Error())
		_ = messager.Send(eventer.StatusError, fmt.Sprintf("failed to setup executor: %s", err.Error()))
		return
	}

	tasks, err := executor.GetTaskList()
	if err != nil {
		logger.Errorf("failed to list tasks: %s", err.Error())
		_ = messager.Send(eventer.StatusError, fmt.Sprintf("failed to list tasks: %s", err.Error()))
		return
	}

	var list []*TaskItem
	for _, taskItem := range tasks {
		list = append(list, toTaskItem(taskItem))
	}
	if len(list) == 0 {
		list = []*TaskItem{}
	}

	_ = messager.Send(eventer.StatusSuccess, list)
}

func (tc *taskClient) createExecute(projectPath, taskFile string, writer io.Writer) *task.Executor {
	return task.NewExecutor(
		task.WithEntrypoint(path.Join(tc.basePath, projectPath, taskFile)),
		task.WithDir(path.Join(tc.basePath, projectPath)),
		task.WithForce(true),
		task.WithVersionCheck(true),
		task.WithColor(false),   // Disable color output for better readability in logs
		task.WithStdout(writer), // Use the provided writer for stdout
		task.WithStderr(writer), // Use the provided writer for stderr
	)
}
