package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
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

	// Create a buffered channel for messages to avoid blocking the task execution
	messageChan := make(chan string, 20)

	// Start to run the project
	pm.taskClient.RunTask(ctx.Request().Context(), project.Path, payload.TaskFile, TaskNameStart, messageChan)

	// Receive messages and send them to the client over server-sent events (SSE)
	options := buildReceiveOptions(
		utils.HeaderUser(ctx),
		project.ID,
		TaskNameStart,
		messageChan,
		func() error {
			return pm.updateProjectStatus(project, StatusRun)
		},
	)

	err := receiveMessageAndSent(options)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}
