package projects

import (
	"pierflow/internal/gitter"
	"strings"
	"time"
)

func toProjectResponse(p *DbProject) *ProjectResponse {
	return &ProjectResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		GitUrl:      p.GitUrl,
		Branch:      p.Branch,
		Path:        p.Path,
		User:        p.User,
		Creation:    time.Unix(p.Creation, 0).Format(time.RFC3339),
		Modified:    time.Unix(p.Modified, 0).Format(time.RFC3339),
		Status:      p.Status.String(),
		CommandMap:  verifier.CommandMap(p.Status),
	}
}

func toProjectMessageListResponse(p *DbProject, message string) *ProjectMessageListResponse {
	return &ProjectMessageListResponse{
		ProjectResponse: *toProjectResponse(p),
		Messages:        strings.Split(message, "\n"),
	}
}

func toProjectBranchMessageListResponse(p *DbProject, branch *gitter.Branch, message string) *ProjectBranchMessageListResponse {
	return &ProjectBranchMessageListResponse{
		ProjectResponse: *toProjectResponse(p),
		Branch:          toBranchInfo(branch),
		Messages:        strings.Split(message, "\n"),
	}
}

func toBranchResponse(branches []gitter.Branch, message string) *BranchListResponse {
	var list []*BranchInfo
	for _, branch := range branches {
		list = append(list, toBranchInfo(&branch))
	}
	if len(list) == 0 {
		list = []*BranchInfo{}
	}
	return &BranchListResponse{
		Branches: list,
		Messages: strings.Split(message, "\n"),
	}
}

func toBranchInfo(b *gitter.Branch) *BranchInfo {
	return &BranchInfo{
		Branch: b.Branch,
		Place:  b.Place.String(),
		Active: b.Active,
	}
}

func toProjectTaskMessageListResponse(p *DbProject, taskFile, taskName, message string) *ProjectTaskMessageListResponse {
	return &ProjectTaskMessageListResponse{
		ProjectResponse: *toProjectResponse(p),
		TaskFile:        taskFile,
		TaskName:        taskName,
		Messages:        strings.Split(message, "\n"),
	}
}
