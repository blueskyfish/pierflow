package projects

import (
	"net/http"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) StopProject(ctx echo.Context) error {
	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandStopProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Stop project '%s' with force", project.Name)
	}

	// Stop the project
	message, err := pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameStop)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to stop project '%s' => %s", project.Name, err.Error()))
	}
	// Update DbProject Status
	err = pm.updateProjectStatus(project, StatusStopped)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to update project status 'stop': %s", err.Error()))
	}
	logger.Infof("Stopped project '%s' with task file '%s'", project.Name, payload.TaskFile)

	// Response
	return ctx.JSON(http.StatusOK, toProjectTaskMessageListResponse(project, payload.TaskFile, TaskNameStop, message))
}
