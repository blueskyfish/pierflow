package projects

import (
	"net/http"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// CloneRepositoryProject try to clone the project into the filesystem with git
func (pm *ProjectManager) CloneRepositoryProject(ctx echo.Context) error {
	projectId := ctx.Param("id")

	var message CommandPayload
	if err := ctx.Bind(&message); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid command '%s' payload", CommandCloneRepository.String()))
	}

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponseF("DbProject with id '%s' not found", projectId))
	}
	logger.Infof("Cloning project '%s' with id '%s'", project.Name, message.Message)

	if err := VerifyCommandToStatus(CommandCloneRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	resultMsg, err := pm.gitClient.Clone(ctx.Request().Context(), project.User, project.Token, project.GitUrl, project.Path)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Cloning is failed in project '%s' => %s", project.Name, err.Error()))
	}

	err = pm.updateProjectStatus(project, StatusCloned)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, toErrorResponse("Failed to update project status into 'cloned'"))
	}

	return ctx.JSON(http.StatusOK, toProjectMessageListResponse(project, resultMsg))
}
