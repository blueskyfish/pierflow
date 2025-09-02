package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/gitter"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

// CloneRepositoryProject try to clone the project into the filesystem with git
func (pm *ProjectManager) CloneRepositoryProject(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

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

	if err := verifier.VerifyStatus(CommandCloneRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	messager := pm.eventServe.WithMessage(CommandCloneRepository.Message(), userId, project.ID, func() {
		if err := pm.updateProjectStatus(project, StatusCloned); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusCloned, err.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("Failed to create messager"))
	}

	// clone the repository
	options := gitter.CloneOptions{
		User:    project.User,
		Token:   project.Token,
		RepoUrl: project.GitUrl,
		Path:    project.Path,
	}
	pm.gitClient.Clone(&options, messager)
	return ctx.String(http.StatusNoContent, "")
}
