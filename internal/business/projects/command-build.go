package projects

import (
	"net/http"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// BuildProject build the project using the specified task file.
//
// The payload TaskFileProjectPayload includes the task file to be used for the build process and a message
// The query parameter "force" can be used to force the build even if the project status is not suitable for building.
func (pm *ProjectManager) BuildProject(ctx echo.Context) error {
	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandBuildProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Build project '%s' with force", project.Name)
	}

	// Build the project
	message, err := pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameBuild)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to build project '%s' => %s", project.Name, err.Error()))
	}

	// Update DbProject Status
	err = pm.updateProjectStatus(project, StatusBuilt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to update project status 'build': %s", err.Error()))
	}
	logger.Infof("Built project '%s' with task file '%s'", project.Name, payload.TaskFile)

	// Response
	return ctx.JSON(http.StatusOK, toProjectTaskMessageListResponse(project, payload.TaskFile, TaskNameBuild, message))
}
