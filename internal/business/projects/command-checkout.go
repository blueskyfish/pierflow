package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
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

	if err := verifier.VerifyStatus(CommandBranchList, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	// build the branch options
	options := gitter.BranchOptions{
		Refresh: refresh,
		Prune:   true,
		Path:    project.Path,
		User:    project.User,
		Token:   project.Token,
	}

	messager := pm.eventServe.WithMessage(CommandBranchList.Message(), userId, project.ID, nil)
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("Failed to create messager"))
	}

	// get the branch list
	logger.Infof("Get branches for project '%s' with refresh=%t", project.Name, refresh)
	pm.gitClient.BranchList(&options, messager)

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

	messager := pm.eventServe.WithMessage(CommandCheckoutRepository.Message(), userId, project.ID, func(data interface{}) {
		branch, ok := data.(gitter.Branch)
		if !ok {
			logger.Warnf("Invalid branch data after checkout for project '%s'", project.Name)
			return
		}
		if err := pm.updateProjectWith(project, StatusCheckedOut, branch.Branch); err != nil {
			logger.Errorf("Failed to update project status to '%s': %s", StatusCheckedOut, err.Error())
		}
	})
	if messager == nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("Failed to create messager"))
	}

	// check out the branch
	options := gitter.CheckoutOptions{
		Branch: payload.Branch,
		Place:  payload.Place,
		Path:   project.Path,
	}

	pm.gitClient.Checkout(&options, messager)
	return ctx.String(http.StatusNoContent, "")
}
