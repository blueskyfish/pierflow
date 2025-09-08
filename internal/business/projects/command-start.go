package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// StartProject starts the project using the specified task file.
//
// The payload TaskFileProjectPayload includes the task file to be used for the start process and a message.
// The query parameter "force" can be used to force the start even if the project status is not suitable for starting.
func (pm *ProjectManager) StartProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	project, force, pErr := pm.prepareProjectTask(ctx, CommandStartProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Start project '%s' with force", project.Name)
	}

	messager := pm.eventServe.WithMessage(CommandStartProject.Message(), userId, project.ID, func(data interface{}) {
		logger.Infof("[%s] finished with %s", project.Name, data.(string))
		if err := pm.updateProjectStatus(project, StatusRun); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusRun, err.Error())
		}
	})

	taskfile := project.Taskfile
	if taskfile == "" {
		taskfile = DefaultTaskfileName
		logger.Infof("Project '%s' does not have a taskfile, using default '%s'", project.Name, taskfile)
	}

	// Start to run the project
	pm.taskClient.Task(project.Path, taskfile, TaskNameStart, messager)

	return ctx.String(http.StatusNoContent, "")
}
