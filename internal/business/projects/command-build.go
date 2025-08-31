package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// BuildProject build the project using the specified task file.
//
// The payload TaskFileProjectPayload includes the task file to be used for the build process and a message
// The query parameter "force" can be used to force the build even if the project status is not suitable for building.
func (pm *ProjectManager) BuildProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User is required")
	}

	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandBuildProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Build project '%s' with force", project.Name)
	}

	messager := pm.eventServe.Messager(userId, CommandBuildProject.Message(), project.ID, func() {
		if err := pm.updateProjectStatus(project, StatusBuilt); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusBuilt, err.Error())
		}
	})

	// Build the project
	pm.taskClient.Task(project.Path, payload.TaskFile, TaskNameBuild, messager)

	return ctx.String(http.StatusNoContent, "")
}
