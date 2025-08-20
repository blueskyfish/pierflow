package projects

import (
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
	// Create a buffered channel for messages to avoid blocking the task execution
	messageChan := make(chan string, 20)

	// Build the project
	pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameBuild, messageChan)

	return receiveMessageAndSent(ctx, messageChan, func() error {
		return pm.updateProjectStatus(project, StatusBuilt)
	})
}
