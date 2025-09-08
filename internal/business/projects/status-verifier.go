package projects

import "fmt"

type statusVerifier struct {
	// listCommand contains all commands that can be used to change the status of a project
	listCommand []ProjectCommand
	// mapStatus maps each command to the valid statuses that can be transitioned to
	// when executing that command.
	// The key is the command, and the value is a slice of valid statuses.
	// This allows us to verify if a command can be executed based on the current status of the project.
	mapStatus map[ProjectCommand][]ProjectStatus
}

// CommandMap is a map that associates ProjectCommand with a boolean value indicating whether the command can be executed for a given status.
type CommandMap map[ProjectCommand]bool

var verifier = statusVerifier{
	listCommand: []ProjectCommand{
		CommandCreateProject,
		CommandCloneRepository,
		CommandCheckoutRepository,
		CommandBuildProject,
		CommandPullRepository,
		CommandStartProject,
		CommandStopProject,
		CommandDeleteProject,
	},
	mapStatus: map[ProjectCommand][]ProjectStatus{
		CommandCloneRepository:    {StatusCreated},
		CommandCheckoutRepository: {StatusCloned, StatusCreated, StatusCheckedOut, StatusBuilt, StatusPulled, StatusStopped},
		CommandBuildProject:       {StatusCheckedOut, StatusBuilt, StatusCloned, StatusPulled, StatusStopped},
		CommandPullRepository:     {StatusCloned, StatusPulled, StatusCheckedOut, StatusStopped, StatusBuilt, StatusCreated}, // TODO check if Created should be here
		CommandStartProject:       {StatusBuilt, StatusStopped},
		CommandStopProject:        {StatusRun},
		CommandDeleteProject:      {StatusStopped, StatusCloned, StatusCheckedOut, StatusBuilt},
	},
}

func (sv *statusVerifier) ListCommand() []ProjectCommand {
	// Return a copy of the list of commands to prevent external modification
	commandsCopy := make([]ProjectCommand, len(sv.listCommand))
	copy(commandsCopy, sv.listCommand)
	return commandsCopy
}

func (sv *statusVerifier) StatusListBy(command ProjectCommand) []ProjectStatus {
	// Return a copy of the valid statuses for the given command
	if command == CommandCreateProject {
		// Create project command does not have a status
		return []ProjectStatus{}
	}
	validStatuses, exists := sv.mapStatus[command]
	if !exists {
		// Command is not recognized, return an empty slice
		return []ProjectStatus{}
	}
	statusesCopy := make([]ProjectStatus, len(validStatuses))
	copy(statusesCopy, validStatuses)
	return statusesCopy
}

func (sv *statusVerifier) VerifyStatus(command ProjectCommand, currentStatus ProjectStatus) error {
	if command == CommandCreateProject || command == CommandBranchList {
		// Create project command does not have a status
		// Branch list command is always allowed
		return nil
	}
	validStatuses, exists := sv.mapStatus[command]
	if !exists {
		// Command is not recognized, return an error
		return fmt.Errorf("unknown command: %s", command)
	}
	for _, status := range validStatuses {
		if status == currentStatus {
			return nil
		}
	}
	return fmt.Errorf("command '%s' is not valid for current status %s", command, currentStatus)
}

// CommandMap returns a map of commands that can be executed based on the current project status.
//
// The keys are the commands, and the values are booleans indicating whether the command can be executed.
// The map is constructed based on the predefined valid statuses for each command.
//
// The command `CommandCreateProject` is excluded from the map since it is always allowed to be executed.
func (sv *statusVerifier) CommandMap(currentStatus ProjectStatus) CommandMap {
	// Create a map to hold the commands that can be executed based on the current status
	commandMap := make(CommandMap)

	for _, command := range sv.listCommand {
		if command == CommandCreateProject {
			continue
		}
		commandMap[command] = false // Initialize all commands to false (not executable)
		validStatuses, exists := sv.mapStatus[command]

		if !exists {
			// not exist, leave it as false
			continue
		}

		for _, status := range validStatuses {
			if status == currentStatus {
				commandMap[command] = true // Command can be executed for the current status
				break                      // No need to check further once we find a match
			}
		}
	}

	return commandMap
}
