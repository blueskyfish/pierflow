package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/gitter"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetProjectBranchPull(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	projectId := ctx.Param("id")

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("DbProject not found"))
	}

	if err := verifier.VerifyStatus(CommandPullRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	messager := pm.eventServe.WithMessage(CommandPullRepository.Message(), userId, project.ID, func() {
		if err := pm.updateProjectStatus(project, StatusPulled); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusPulled, err.Error())
		}
	})

	// pull the repository
	options := gitter.PullOptions{
		User:   project.User,
		Token:  project.Token,
		GitUrl: project.GitUrl,
		Path:   project.Path,
	}
	pm.gitClient.Pull(ctx.Request().Context(), &options, messager)

	return ctx.String(http.StatusNoContent, "")
}
