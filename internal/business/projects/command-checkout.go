package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/eventer"
	"pierflow/internal/gitter"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetProjectBranchList(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	projectId := ctx.Param("id")
	refresh := utils.QueryBool(ctx, "refresh", false)

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse("Not found project"))
	}

	if err := verifier.VerifyStatus(CommandCreateProject, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	// build the branch options
	options := gitter.BranchOptions{
		Refresh: refresh,
	}
	if refresh {
		options.User = project.User
		options.Token = project.Token
		options.Prune = true
		options.Path = project.Path
	}

	messager := eventer.NewMessager(eventer.StatusDebug, nil)

	logger.Infof("Get branches for project '%s' with refresh=%t", project.Name, refresh)
	pm.gitClient.BranchList(ctx.Request().Context(), &options, messager)

	// wait for the operation to complete
	err := pm.listenEventMessager(userId, project.ID, CommandCreateProject.String(), messager, nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to get branches for project '%s' => %s", project.Name, err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}

func (pm *ProjectManager) CheckoutProjectBranch(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("User header is required"))
	}

	projectId := ctx.Param("id")
	var payload CheckoutPayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid payload for project '%s'", projectId))
	}

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse("Not found project"))
	}
	if err := verifier.VerifyStatus(CommandCheckoutRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}
	logger.Infof("Checkout project '%s' branch '%s'", project.Name, payload.Branch)

	messager := eventer.NewMessager(eventer.StatusDebug, nil)

	options := gitter.CheckoutOptions{
		Branch: payload.Branch,
		Place:  payload.Place,
		Path:   project.Path,
	}

	pm.gitClient.Checkout(&options, messager)

	err := pm.listenEventMessager(userId, project.ID, CommandCheckoutRepository.String(), messager, func() error {
		return pm.updateProjectStatus(project, StatusCheckedOut)
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Checkout is failed in project '%s' => %s", project.Name, err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}
