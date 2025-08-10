package projects

import (
	"net/http"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// StartProject starts the project using the specified task file.
//
// The payload TaskFileProjectPayload includes the task file to be used for the start process and a message.
// The query parameter "force" can be used to force the start even if the project status is not suitable for starting.
func (pm *ProjectManager) StartProject(ctx echo.Context) error {
	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandStartProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Start project '%s' with force", project.Name)
	}

	// Start to run the project
	message, err := pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameStart)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to start project '%s' => %s", project.Name, err.Error()))
	}

	// Update DbProject Status
	err = pm.updateProjectStatus(project, StatusRun)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to update project status 'run': %s", err.Error()))
	}
	logger.Infof("Started project '%s' with task file '%s'", project.Name, payload.TaskFile)

	// Response
	return ctx.JSON(http.StatusOK, toProjectTaskMessageListResponse(project, payload.TaskFile, TaskNameStart, message))
}
