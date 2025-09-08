package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) StopProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	project, force, pErr := pm.prepareProjectTask(ctx, CommandStopProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Stop project '%s' with force", project.Name)
	}

	messager := pm.eventServe.WithMessage(CommandStopProject.Message(), userId, project.ID, func(data interface{}) {
		logger.Infof("[%s] finished with %s", project.Name, data.(string))
		if err := pm.updateProjectStatus(project, StatusStopped); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusStopped, err.Error())
		}
	})

	taskfile := project.Taskfile
	if taskfile == "" {
		taskfile = DefaultTaskfileName
		logger.Infof("Project '%s' does not have a taskfile, using default '%s'", project.Name, taskfile)
	}

	// Stop the project
	pm.taskClient.Task(project.Path, taskfile, TaskNameStop, messager)

	return ctx.String(http.StatusNoContent, "")
}
