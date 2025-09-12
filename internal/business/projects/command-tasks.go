package projects

import (
	"github.com/blueskyfish/pierflow/internal/business/utils"
	"github.com/blueskyfish/pierflow/internal/errors"

	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (pm *ProjectManager) GetTaskFileList(ctx echo.Context) error {
	projectId := ctx.Param("id")

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, errors.ToErrorResponse("Not found project"))
	}

	taskFiles, err := pm.listTaskFiles(project)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponseF("Failed to list task files for project '%s' => %s", project.Name, err.Error()))
	}

	return ctx.JSON(http.StatusOK, taskFiles)
}

// UpdateTaskfile updates the taskfile of a project.
//
// The payload ChangeTaskfilePayload includes the new taskfile to be set for the project.
func (pm *ProjectManager) UpdateTaskfile(ctx echo.Context) error {
	projectId := ctx.Param("id")

	var payload ChangeTaskfilePayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Invalid payload"))
	}

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, errors.ToErrorResponse("Not found project"))
	}

	err := pm.db.Transaction(func(tx *gorm.DB) error {
		project.Taskfile = payload.Taskfile
		return tx.Save(project).Error
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Failed to update taskfile"))
	}

	return ctx.JSON(http.StatusOK, toProjectResponse(project))
}

// GetTaskNameList retrieves the list of task names from the task file of a project.
func (pm *ProjectManager) GetTaskNameList(ctx echo.Context) error {
	userId := utils.HeaderUser(ctx)
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("User header is required"))
	}

	projectId := ctx.Param("id")
	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, errors.ToErrorResponse("Not found project"))
	}

	taskfile := project.Taskfile
	if taskfile == "" {
		taskfile = "Taskfile.yml"
	}

	messager := pm.eventServe.WithMessage(CommandTaskList.Message(), userId, project.ID, nil)

	// Start to list the task names
	pm.taskClient.List(project.Path, taskfile, messager)

	return ctx.String(http.StatusNoContent, "")
}

// prepareProjectTask prepares the project for a specific command task.
//
// It retrieves the project by ID from the context, checks if the command can be executed based on the project's status,
// and returns the project, a force flag, and any error encountered during the process.
func (pm *ProjectManager) prepareProjectTask(ctx echo.Context, command ProjectCommand) (*DbProject, bool, *ProjectError) {
	projectId := ctx.Param("id")
	force := utils.QueryBool(ctx, "force", false)

	// DbProject loading
	project := pm.findProjectById(projectId)
	if project == nil {
		return nil, false, toError(http.StatusNotFound, "DbProject not found")
	}

	if !force {
		err := verifier.VerifyStatus(command, project.Status)
		if err != nil {
			return nil, false, toErrorF(http.StatusBadRequest, "Invalid project status '%s' to run project", project.Status)
		}
	}

	return project, force, nil
}
