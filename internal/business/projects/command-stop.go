package projects

import (
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

	// Create a buffered channel for messages to avoid blocking the task execution
	messageChan := make(chan string, 20)

	// Stop the project
	pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameStop, messageChan)

	// Response
	return receiveMessageAndSent(ctx, messageChan, func() error {
		return pm.updateProjectStatus(project, StatusStopped)
	})
}
