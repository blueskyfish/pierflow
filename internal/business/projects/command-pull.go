package projects

import (
	"github.com/blueskyfish/pierflow/internal/business/errors"
	"github.com/blueskyfish/pierflow/internal/business/utils"
	"github.com/blueskyfish/pierflow/internal/gitter"
	"github.com/blueskyfish/pierflow/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetProjectBranchPull(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("User header is required"))
	}

	projectId := ctx.Param("id")

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("DbProject not found"))
	}

	if err := verifier.VerifyStatus(CommandPullRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	messager := pm.eventServe.WithMessage(CommandPullRepository.Message(), userId, project.ID, func(data interface{}) {
		var result gitter.Branch
		err := utils.CovertTo(data, &result)
		if err != nil {
			logger.Errorf("Pull the project %s failed: %v", project.Name, err)
			return
		}
		logger.Infof("[%s] finished with branch %s", project.Name, result.Branch)
		// update the project branch and status
		if err := pm.updateProjectWith(project, StatusPulled, result.Branch, CommandPullRepository.Event()); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusPulled, err.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Failed to create messager"))
	}

	// pull the repository
	options := gitter.PullOptions{
		User:   project.User,
		Token:  project.Token,
		GitUrl: project.GitUrl,
		Path:   project.Path,
	}
	pm.gitClient.Pull(&options, messager)

	return ctx.String(http.StatusNoContent, "")
}
