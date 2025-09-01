package projects

import (
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
