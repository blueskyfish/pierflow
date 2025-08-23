package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
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

	// Create an unbuffered channel for messages to avoid blocking the task execution
	messageChan := make(chan string)

	// Stop the project
	pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameStop, messageChan)

	// Receive messages and send them to the client over server-sent events (SSE)
	options := buildReceiveOptions(
		utils.HeaderUser(ctx),
		project.ID,
		TaskNameBuild,
		messageChan,
		func() error {
			return pm.updateProjectStatus(project, StatusStopped)
		},
	)

	err := receiveMessageAndSent(options)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}
