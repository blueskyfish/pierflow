package projects

import (
	"encoding/json"
	"time"
)

func toProjectResponse(p *DbProject) *ProjectResponse {
	return &ProjectResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		GitUrl:      p.GitUrl,
		Branch:      p.Branch,
		Taskfile:    p.Taskfile,
		Path:        p.Path,
		User:        p.User,
		Creation:    time.Unix(p.Creation, 0).Format(time.RFC3339),
		Modified:    time.Unix(p.Modified, 0).Format(time.RFC3339),
		Status:      p.Status.String(),
		CommandMap:  verifier.CommandMap(p.Status),
	}
}

func toProjectHistoryResponse(ev *DbEvent) *ProjectHistoryResponse {
	var project DbProject
	err := json.Unmarshal([]byte(ev.Value), &project)
	if err != nil {
		// If we can't unmarshal the project, return a minimal event response
		return &ProjectHistoryResponse{
			ID:        ev.ID,
			Group:     ev.Group,
			Event:     ev.Event,
			Project:   &ProjectResponse{ID: ev.Value}, // Just return the raw value as ID
			Timestamp: ev.Timestamp,
		}
	}

	return &ProjectHistoryResponse{
		ID:        ev.ID,
		Group:     ev.Group,
		Event:     ev.Event,
		Project:   toProjectResponse(&project),
		Timestamp: ev.Timestamp,
	}
}
