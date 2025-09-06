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
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User is required"))
	}

	project, payload, force, err := pm.prepareProjectTask(ctx, CommandBuildProject)
	if err != nil {
		return err.JSON(ctx)
	}

	if force {
		logger.Infof("Build project '%s' with force", project.Name)
	}

	messager := pm.eventServe.WithMessage(CommandBuildProject.Message(), userId, project.ID, func(data interface{}) {
		logger.Infof("[%s] finished with %s", project.Name, data.(string))
		if pErr := pm.updateProjectStatus(project, StatusBuilt); pErr != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusBuilt, pErr.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("Failed to create messager"))
	}

	// Build the project
	pm.taskClient.Task(project.Path, payload.TaskFile, TaskNameBuild, messager)

	return ctx.String(http.StatusNoContent, "")
}
