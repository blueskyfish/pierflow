package api

import (
	"errors"
	"pierflow/internal/business"
	"pierflow/internal/business/projects"

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

	// Manages DbProject Commands
	group.PUT("/projects/:id/clone", pm.CloneRepositoryProject)
	group.PUT("/projects/:id/checkout", pm.CheckoutProjectBranch)
	group.PUT("/projects/:id/build", pm.BuildProject)
	group.PUT("/projects/:id/start", pm.StartProject)
	group.PUT("/projects/:id/stop", pm.StopProject)

	// Manage DbProject Branches
	group.GET("/projects/:id/branches", pm.GetProjectBranchList)
	group.GET("/projects/:id/branches/pull", pm.GetProjectBranchPull)

	group.GET("/projects/:id/tasks", pm.GetTaskFileList)
	group.GET("/projects/:id/tasks/:taskFile", pm.GetTaskNameListByTaskFile)

	// Managed System
	group.GET("/system", sm.GetSystemInfo)

	// User Management
	group.GET("/users/:id", pm.UserConnect)
	group.GET("/users/:id/ping", pm.SendPing)

	return nil
}
