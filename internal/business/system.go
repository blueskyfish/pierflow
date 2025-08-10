package business

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
)

type SystemManager struct {
	basePath string
}

type CommandInfo struct {
	Command string `json:"command"`
	Version string `json:"version,omitempty"`
	Path    string `json:"path,omitempty"`
	Error   string `json:"error,omitempty"`
}

type SystemInfoResponse struct {
	Commands []CommandInfo `json:"commands"`
}

func NewSystemManager(basePath string) *SystemManager {
	return &SystemManager{
		basePath: basePath,
	}
}

func (sm *SystemManager) GetSystemInfo(ctx echo.Context) error {
	commands := []string{"git", "docker", "task"}

	var commandList []CommandInfo
	for _, command := range commands {
		cmdInfo := CommandInfo{Command: command}
		path, err := sm.findCommand(command)
		if err != nil {
			cmdInfo.Error = err.Error()
		} else {
			cmdInfo.Path = path
			version, err := sm.findCommandVersion(command)
			if err != nil {
				cmdInfo.Error = err.Error()
			} else {
				cmdInfo.Version = version
			}
		}
		commandList = append(commandList, cmdInfo)
	}

	return ctx.JSON(http.StatusOK, &SystemInfoResponse{
		Commands: commandList,
	})
}

func (sm *SystemManager) findCommand(command string) (string, error) {
	out, err := exec.Command("which", command).Output()
	if err != nil {
		return "", err
	}
	path := strings.TrimSpace(string(out))
	if path == "" {
		return "", exec.ErrNotFound
	}
	return path, nil
}

func (sm *SystemManager) findCommandVersion(command string) (string, error) {
	out, err := exec.Command(command, "--version").Output()
	if err != nil {
		return "", err
	}
	version := strings.TrimSpace(string(out))
	if version == "" {
		return "", exec.ErrNotFound
	}
	return version, nil
}
