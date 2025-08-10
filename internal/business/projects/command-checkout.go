package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/gitter"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetProjectBranchList(ctx echo.Context) error {
	projectId := ctx.Param("id")
	refresh := utils.QueryBool(ctx, "refresh", false)

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(404, toErrorResponse("Not found project"))
	}

	if err := VerifyCommandToStatus(CommandCreateProject, project.Status); err != nil {
		return ctx.JSON(400, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}

	// build the branch options
	options := &gitter.BranchOptions{
		Refresh: refresh,
	}
	if refresh {
		options.User = project.User
		options.Token = project.Token
		options.Prune = true
	}

	message, branches, err := pm.gitClient.BranchList(ctx.Request().Context(), project.Path, options)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to get branches for project '%s' => %s", project.Name, err.Error()))
	}

	return ctx.JSON(http.StatusOK, toBranchResponse(branches, message))
}

func (pm *ProjectManager) CheckoutProjectBranch(ctx echo.Context) error {
	projectId := ctx.Param("id")
	var payload CheckoutPayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid payload for project '%s'", projectId))
	}

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse("Not found project"))
	}
	if err := VerifyCommandToStatus(CommandCheckoutRepository, project.Status); err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Invalid project status %s => %s", project.Status, err.Error()))
	}
	logger.Infof("Checkout project '%s' branch '%s'", project.Name, payload.Branch)

	options := &gitter.CheckoutOptions{
		Branch: payload.Branch,
		Place:  payload.Place,
	}

	branch, err := pm.gitClient.CheckoutBranch(project.Path, options)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to checkout branch '%s' in project '%s' => %s", payload.Branch, project.Name, err.Error()))
	}

	err = pm.updateProjectStatus(project, StatusCheckedOut)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, toErrorResponse("Failed to update project status to 'checked out'"))
	}

	logger.Infof("DbProject '%s' branch '%s' is checked out successfully", project.Name, branch.Branch)
	return ctx.JSON(http.StatusOK, toBranchInfo(branch))
}
