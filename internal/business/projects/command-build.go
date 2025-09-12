package projects

import (
	"net/http"
	"pierflow/internal/business/errors"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// BuildProject build the project using the specified task file.
//
// The query parameter "force" can be used to force the build even if the project status is not suitable for building.
func (pm *ProjectManager) BuildProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("User is required"))
	}

	project, force, err := pm.prepareProjectTask(ctx, CommandBuildProject)
	if err != nil {
		return err.JSON(ctx)
	}

	if force {
		logger.Infof("Build project '%s' with force", project.Name)
	}

	messager := pm.eventServe.WithMessage(CommandBuildProject.Message(), userId, project.ID, func(data interface{}) {
		logger.Infof("[%s] finished with %s", project.Name, data.(string))
		if pErr := pm.updateProjectStatus(project, StatusBuilt, CommandBuildProject.Event()); pErr != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusBuilt, pErr.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Failed to create messager"))
	}

	taskfile := project.Taskfile
	if taskfile == "" {
		taskfile = DefaultTaskfileName
		logger.Infof("Project '%s' does not have a taskfile, using default '%s'", project.Name, taskfile)
	}

	// Build the project
	pm.taskClient.Task(project.Path, taskfile, TaskNameBuild, messager)

	return ctx.String(http.StatusNoContent, "")
}
