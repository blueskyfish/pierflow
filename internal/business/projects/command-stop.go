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

	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandStopProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Stop project '%s' with force", project.Name)
	}

	messager := pm.eventServe.WithMessage(CommandStopProject.Message(), userId, project.ID, func() {
		if err := pm.updateProjectStatus(project, StatusStopped); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusStopped, err.Error())
		}
	})

	// Stop the project
	pm.taskClient.Task(project.Path, payload.TaskFile, TaskNameStop, messager)

	return ctx.String(http.StatusNoContent, "")
}
