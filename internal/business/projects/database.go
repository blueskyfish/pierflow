package projects

import (
	"net/http"
	"pierflow/internal/business/errors"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (pm *ProjectManager) GetProjectList(ctx echo.Context) error {
	var projects []DbProject
	if err := pm.db.Find(&projects).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.ToErrorResponse("Failed to retrieve projects"))
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
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponseF("DbProject payload is invalid"))
	}
	now := time.Now().UTC()

	project := DbProject{
		ID:          "",
		Name:        payload.Name,
		Description: payload.Description,
		GitUrl:      payload.GitUrl,
		Path:        payload.Path,
		Branch:      "",
		User:        payload.User,
		Token:       payload.Token,
		Creation:    now.Unix(),
		Modified:    now.Unix(),
		Status:      StatusCreated,
	}

	err := pm.db.Transaction(func(tx *gorm.DB) error {
		// Create the project
		err := tx.Create(&project).Error
		if err != nil {
			return err
		}
		// Create an event for the project creation
		return pm.createEvent(tx, &project, CommandCreateProject.Event())
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponse("Failed to create project"))
	}

	return ctx.JSON(http.StatusCreated, toProjectResponse(&project))
}

func (pm *ProjectManager) GetProjectDetail(ctx echo.Context) error {
	projectId := ctx.Param("id")
	if projectId == "" {
		return ctx.JSON(http.StatusBadRequest, errors.ToErrorResponseF("Project ID is required"))
	}

	project := pm.findProjectById(projectId)
	if project == nil {
		return ctx.JSON(http.StatusNotFound, errors.ToErrorResponse("Project not found"))
	}

	return ctx.JSON(http.StatusOK, toProjectResponse(project))
}

func (pm *ProjectManager) findProjectById(projectId string) *DbProject {
	var project DbProject
	if err := pm.db.Find(&project, "id = ?", projectId).Error; err != nil {
		logger.Warnf("DbProject with ID '%s' not found: %s", projectId, err.Error())
		return nil
	}
	return &project
}

func (pm *ProjectManager) updateProjectStatus(p *DbProject, status ProjectStatus, eventName string) error {
	return pm.db.Transaction(func(tx *gorm.DB) error {
		p.Status = status
		p.Modified = tx.NowFunc().Unix()
		err := tx.Save(p).Error
		if err != nil {
			logger.Warnf("DbProject with ID '%s' not found: %s", p.ID, err.Error())
			return err
		}

		// Create an event for the project status update
		return pm.createEvent(tx, p, eventName)
	})
}

func (pm *ProjectManager) updateProjectWith(p *DbProject, status ProjectStatus, branch, eventName string) error {
	return pm.db.Transaction(func(tx *gorm.DB) error {
		// Update the project status and branch
		p.Status = status
		p.Branch = branch
		p.Modified = tx.NowFunc().Unix()
		err := tx.Save(p).Error
		if err != nil {
			logger.Warnf("DbProject with ID '%s' not found: %s", p.ID, err.Error())
			return err
		}
		// Create an event for the project update
		return pm.createEvent(tx, p, eventName)
	})
}

func (pm *ProjectManager) createEvent(tx *gorm.DB, p *DbProject, eventName string) error {

	// JSON serialize the project
	value, err := utils.Stringify(p)
	if err != nil {
		logger.Warnf("DbProject with ID '%s' is invalid: %s", p.ID, err.Error())
		return err
	}

	//
	dbEvent := DbEvent{
		ID:        "",
		Group:     "project",
		Event:     eventName,
		ValueID:   p.ID,
		Value:     value,
		Timestamp: tx.NowFunc().Unix(),
	}
	return tx.Create(&dbEvent).Error
}
