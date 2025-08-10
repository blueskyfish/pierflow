package projects

type ProjectCommand string

const (
	CommandCreateProject      ProjectCommand = "create-project"
	CommandCloneRepository    ProjectCommand = "clone-repository"
	CommandCheckoutRepository ProjectCommand = "checkout-repository"
	CommandBuildProject       ProjectCommand = "build-project"
	CommandPullRepository     ProjectCommand = "pull-repository" // Update repository to the latest version
	CommandStartProject       ProjectCommand = "start-project"
	CommandStopProject        ProjectCommand = "stop-project"
	CommandDeleteProject      ProjectCommand = "delete-project"
	CommandTaggingVersion     ProjectCommand = "tagging-version"
)

func (c ProjectCommand) String() string {
	return string(c)
}

func ProjectCommandFrom(s string) (ProjectCommand, bool) {
	switch s {
	case CommandCreateProject.String(),
		CommandCloneRepository.String(),
		CommandCheckoutRepository.String(),
		CommandBuildProject.String(),
		CommandPullRepository.String(),
		CommandStartProject.String(),
		CommandStopProject.String(),
		CommandTaggingVersion.String(),
		CommandDeleteProject.String(),
		CommandTaggingVersion.String():
		return ProjectCommand(s), true
	default:
		return "", false
	}
}
