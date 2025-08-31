package projects

import (
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func toErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

func toErrorResponseF(format string, args ...any) *ErrorResponse {
	return toErrorResponse(fmt.Sprintf(format, args...))
}

type ProjectResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Path        string     `json:"path"`
	GitUrl      string     `json:"gitUrl"`
	Branch      string     `json:"branch"`
	User        string     `json:"user"`
	Creation    string     `json:"creation"`
	Modified    string     `json:"modified"`
	Status      string     `json:"status"`
	CommandMap  CommandMap `json:"commandMap,omitempty"`
}

type ProjectMessageListResponse struct {
	ProjectResponse
	Messages []string `json:"messages"`
}

type ProjectBranchMessageListResponse struct {
	ProjectResponse
	Branch   *BranchInfo `json:"branch"`
	Messages []string    `json:"messages"`
}

type ProjectTaskMessageListResponse struct {
	ProjectResponse
	TaskFile string   `json:"taskfile"`
	TaskName string   `json:"taskname"`
	Messages []string `json:"messages"`
}

type BranchInfo struct {
	Branch string `json:"branch"`
	Place  string `json:"place"`
	Active bool   `json:"active"`
}

type BranchListResponse struct {
	Branches []*BranchInfo `json:"branches"`
	Messages []string      `json:"messages"`
}
