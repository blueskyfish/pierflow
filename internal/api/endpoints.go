package api

import (
	"errors"

	"github.com/blueskyfish/pierflow/internal/business"
	"github.com/blueskyfish/pierflow/internal/business/projects"

	"github.com/labstack/echo/v4"
)

func registerEndpoints(pm *projects.ProjectManager, sm *business.SystemManager, group *echo.Group) error {
	if group == nil {
		return errors.New("echo group is required")
	}
	if pm == nil {
		return errors.New("project manager is required")
	}
	if sm == nil {
		return errors.New("system manager is required")
	}

	// Manage Projects
	group.GET("/projects", pm.GetProjectList)
	group.POST("/projects", pm.CreateProject)
	group.GET("/projects/:id", pm.GetProjectDetail)
	group.GET("/projects/:id/history", pm.GetProjectHistory)

	// Manages Project Commands
	group.PUT("/projects/:id/clone", pm.CloneRepositoryProject)
	group.PUT("/projects/:id/checkout", pm.CheckoutProjectBranch)
	group.GET("/projects/:id/build", pm.BuildProject)
	group.GET("/projects/:id/start", pm.StartProject)
	group.GET("/projects/:id/stop", pm.StopProject)

	// Manage Project Branches
	group.GET("/projects/:id/branches", pm.GetProjectBranchList)
	group.GET("/projects/:id/branches/pull", pm.GetProjectBranchPull)

	group.GET("/projects/:id/tasks", pm.GetTaskFileList)
	group.PUT("/projects/:id/tasks", pm.UpdateTaskfile)
	group.GET("/projects/:id/tasks/names", pm.GetTaskNameList)

	// Project Event (server-sent events) and connection management
	group.GET("/projects/connect/:id", pm.ProjectEventConnect)
	group.GET("/projects/ping/:id", pm.ProjectEventPing)

	// Managed System
	group.GET("/system", sm.GetSystemInfo)

	return nil
}
