package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/eventer"
	"pierflow/internal/gitter"

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

	messager := eventer.NewMessager(eventer.StatusDebug, nil)

	options := gitter.PullOptions{
		User:   project.User,
		Token:  project.Token,
		GitUrl: project.GitUrl,
		Path:   project.Path,
	}

	pm.gitClient.Pull(ctx.Request().Context(), &options, messager)
	err := pm.listenEventMessager(userId, project.ID, CommandPullRepository.String(), messager, func() error {
		return pm.updateProjectStatus(project, StatusPulled)
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Pulling is failed in project '%s' => %s", project.Name, err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}
