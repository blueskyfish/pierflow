package projects

import "fmt"

type ProjectCommand string

const (
	// CommandCreateProject create a new project
	// It is the first command to be executed when creating a project.
	// It is always allowed to be executed, regardless of the current status of the project.
	CommandCreateProject      ProjectCommand = "create-project"
	CommandCloneRepository    ProjectCommand = "clone-repository"
	CommandCheckoutRepository ProjectCommand = "checkout-repository"
	CommandBuildProject       ProjectCommand = "build-project"
	CommandPullRepository     ProjectCommand = "pull-repository" // Update repository to the latest version
	CommandStartProject       ProjectCommand = "start-project"
	CommandStopProject        ProjectCommand = "stop-project"
	CommandDeleteProject      ProjectCommand = "delete-project"
	CommandTaggingVersion     ProjectCommand = "tagging-version"

	CommandTaskList   ProjectCommand = "task-list"   // Not directly a command, but used to list tasks in a task file
	CommandBranchList ProjectCommand = "branch-list" // Not directly a command, but used to list branches in a repository
)

func (c ProjectCommand) String() string {
	return string(c)
}

func (c ProjectCommand) Message() string {
	return fmt.Sprintf("message-%s", c.String())
}

func (c ProjectCommand) Event() string {
	return fmt.Sprintf("event-%s", c.String())
}
