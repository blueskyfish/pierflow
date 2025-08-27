package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/eventer"
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

	messager := eventer.NewMessager(eventer.StatusDebug, nil)

	// Stop the project
	pm.taskClient.RunTask(project.Path, payload.TaskFile, TaskNameStop, messager)

	err := pm.listenEventMessager(userId, project.ID, CommandStopProject.String(), messager, func() error {
		return pm.updateProjectStatus(project, StatusStopped)
	})
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}
