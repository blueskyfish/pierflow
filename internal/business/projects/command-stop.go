package projects

import (
	"github.com/blueskyfish/pierflow/internal/business/errors"
	"github.com/blueskyfish/pierflow/internal/business/utils"
	"github.com/blueskyfish/pierflow/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) StopProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("User header is required"))
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
		if err := pm.updateProjectStatus(project, StatusStopped, CommandStopProject.Event()); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusStopped, err.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Failed to create messager"))
	}

	taskfile := project.Taskfile
	if taskfile == "" {
		taskfile = DefaultTaskfileName
		logger.Infof("Project '%s' does not have a taskfile, using default '%s'", project.Name, taskfile)
	}

	// Stop the project
	pm.taskClient.Task(project.Path, taskfile, TaskNameStop, messager)

	return ctx.String(http.StatusNoContent, "")
}
