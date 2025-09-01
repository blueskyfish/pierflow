package projects

import (
	"net/http"
	"pierflow/internal/logger"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (pm *ProjectManager) GetProjectList(ctx echo.Context) error {
	var projects []DbProject
	if err := pm.db.Find(&projects).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, &ErrorResponse{Message: "Failed to retrieve projects"})
	}

	var list []*ProjectResponse
	for _, project := range projects {
		list = append(list, toProjectResponse(&project))
	}

	if len(list) == 0 {
		list = []*ProjectResponse{}
	}

	return ctx.JSON(http.StatusOK, list)
}

func (pm *ProjectManager) CreateProject(ctx echo.Context) error {
	var payload CreatePayload
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, &ErrorResponse{Message: "DbProject payload is invalid"})
	}
	now := time.Now()

	project := DbProject{
		ID:          "",
		Name:        payload.Name,
		Description: payload.Description,
		GitUrl:      payload.GitUrl,
		Path:        payload.Path,
		Branch:      "",
		Creation:    now.UTC().Unix(),
		Modified:    now.UTC().Unix(),
		Status:      StatusCreated,
	}

	err := pm.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&project).Error
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &ErrorResponse{Message: "Failed to create project"})
	}

	return ctx.JSON(http.StatusCreated, toProjectResponse(&project))
}

func (pm *ProjectManager) findProjectById(projectId string) *DbProject {
	var project DbProject
	if err := pm.db.Find(&project, "id = ?", projectId).Error; err != nil {
		logger.Warnf("DbProject with ID '%s' not found: %s", projectId, err.Error())
		return nil
	}
	return &project
}

func (pm *ProjectManager) updateProjectStatus(p *DbProject, status ProjectStatus) error {
	return pm.db.Transaction(func(tx *gorm.DB) error {
		p.Status = status
		p.Modified = time.Now().UTC().Unix()
		return tx.Save(p).Error
	})
}
