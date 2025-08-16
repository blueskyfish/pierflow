package projects

import (
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) GetTaskFileList(ctx echo.Context) error {
	projectId := ctx.Param("id")

	var project DbProject
	if err := pm.db.Find(&project, "id = ?", projectId).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse("Not found project"))
	}

	taskFiles, err := pm.listTaskFiles(&project)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, toErrorResponseF("Failed to list task files for project '%s' => %s", project.Name, err.Error()))
	}

	return ctx.JSON(http.StatusOK, taskFiles)
}

func (pm *ProjectManager) GetTaskNameListByTaskFile(ctx echo.Context) error {
	projectId := ctx.Param("id")
	taskFile := ctx.Param("taskFile")

	var project DbProject
	if err := pm.db.Find(&project, "id = ?", projectId).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse("Not found project"))
	}

	if taskFile == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("Task file name is required"))
	}

	taskNames, err := pm.taskClient.TaskList(project.Path, taskFile)
	if err != nil {
		logger.Warnf("Task file '%s' not found in project '%s': %s", taskFile, project.Name, err.Error())
		return ctx.JSON(http.StatusNotFound, toErrorResponseF("Taskfile '%s' not found", taskFile))
	}

	return ctx.JSON(http.StatusOK, taskNames)
}

func (pm *ProjectManager) prepareProjectTask(ctx echo.Context, command ProjectCommand) (*DbProject, *TaskFileProjectPayload, bool, *ProjectError) {
	projectId := ctx.Param("id")
	force := utils.QueryBool(ctx, "force", false)

	// Payload
	var payload TaskFileProjectPayload
	if err := ctx.Bind(&payload); err != nil {
		return nil, nil, false, toErrorF(http.StatusBadRequest, "Invalid payload for project '%s'", projectId)
	}

	// DbProject loading
	project := pm.findProjectById(projectId)
	if project == nil {
		return nil, nil, false, toError(http.StatusNotFound, "DbProject not found")
	}

	if !force {
		err := verifier.VerifyStatus(command, project.Status)
		if err != nil {
			return nil, nil, false, toErrorF(http.StatusBadRequest, "Invalid project status '%s' to run project", project.Status)
		}
	}

	return project, &payload, force, nil
}
