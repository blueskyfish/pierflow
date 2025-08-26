package projects

import (
	"net/http"
	"pierflow/internal/gitter"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetProjectBranchPull(ctx echo.Context) error {
	projectId := ctx.Param("id")

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("DbProject not found"))
	}

	if err := verifier.VerifyStatus(CommandPullRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	options := &gitter.PullOptions{
		User:   project.User,
		Token:  project.Token,
		GitUrl: project.GitUrl,
	}

	resultMsg, branch, err := pm.gitClient.Pull(ctx.Request().Context(), project.Path, options)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Pulling is failed in project '%s' => %s", project.Name, err.Error()))
	}

	return ctx.JSON(http.StatusOK, toProjectBranchMessageListResponse(project, branch, resultMsg))
}
