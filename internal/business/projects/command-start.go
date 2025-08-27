package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// StartProject starts the project using the specified task file.
//
// The payload TaskFileProjectPayload includes the task file to be used for the start process and a message.
// The query parameter "force" can be used to force the start even if the project status is not suitable for starting.
func (pm *ProjectManager) StartProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	project, payload, force, pErr := pm.prepareProjectTask(ctx, CommandStartProject)
	if pErr != nil {
		return pErr.JSON(ctx)
	}

	if force {
		logger.Infof("Start project '%s' with force", project.Name)
	}

	messager := eventer.NewMessager(eventer.StatusDebug, nil)

	// Start to run the project
	pm.taskClient.RunTask(project.Path, payload.TaskFile, TaskNameStart, messager)

	err := pm.listenEventMessager(userId, project.ID, CommandStartProject.String(), messager, func() error {
		return pm.updateProjectStatus(project, StatusRun)
	})

	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}
