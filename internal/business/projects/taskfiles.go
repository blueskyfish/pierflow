package projects

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/blueskyfish/pierflow/internal/errors"
)

const (
	TaskNameList  = "list"  // TaskNameList is used to list all tasks in the taskfile
	TaskNameBuild = "build" // TaskNameBuild is used to build the project
	TaskNameStart = "start" // TaskNameStart is used to start the task
	TaskNameStop  = "stop"  // TaskNameStop is used to stop the task
	TaskNameInfo  = "info"  // TaskNameInfo is used to get information about tasks

	DefaultTaskfileName = "Taskfile.yml"
)

func (pm *ProjectManager) listTaskFiles(p *DbProject) ([]string, error) {
	if p == nil {
		return nil, errors.NewFromText("project is nil")
	}

	var fileList []string
	err := filepath.Walk(path.Join(pm.basePath, p.Path), func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Check if the file is a YAML file and contains "taskfile" in its name (case-insensitive)
		// We consider both .yml and .yaml extensions
		// Example valid names: Taskfile.yml, taskfile.yaml, MyTaskFile.YML
		// Example invalid names: config.yml, tasks.yaml, Taskfile.txt
		if (filepath.Ext(filePath) == ".yml" || filepath.Ext(filePath) == ".yaml") &&
			strings.Contains(strings.ToLower(filepath.Base(filePath)), "taskfile") {
			fileList = append(fileList, info.Name())
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(fileList) == 0 {
		fileList = []string{}
	}
	return fileList, nil
}
