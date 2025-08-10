package tasker

import "github.com/go-task/task/v3/taskfile/ast"

func toTaskItem(task *ast.Task) *TaskItem {
	return &TaskItem{
		Name:    task.Name(),
		Desc:    task.Desc,
		Summary: task.Summary,
	}
}
