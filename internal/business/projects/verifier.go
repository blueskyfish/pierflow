package projects

import (
	"fmt"
)

var mapCommandToStatus = map[ProjectCommand][]ProjectStatus{
	CommandCloneRepository:    {StatusCreated},
	CommandCheckoutRepository: {StatusCloned, StatusCreated, StatusCheckedOut, StatusStopped},
	CommandBuildProject:       {StatusCheckedOut, StatusBuilt, StatusCloned, StatusPulled},
	CommandPullRepository:     {StatusCloned, StatusPulled, StatusCheckedOut, StatusStopped, StatusBuilt, StatusCreated},
	CommandStartProject:       {StatusBuilt, StatusStopped},
	CommandStopProject:        {StatusRun},
	CommandDeleteProject:      {StatusStopped, StatusCloned, StatusCheckedOut, StatusBuilt},
}

func VerifyCommandToStatus(command ProjectCommand, currentStatus ProjectStatus) error {
	if command == CommandCreateProject {
		// Create project command does not have a status
		return nil
	}
	validStatuses, exists := mapCommandToStatus[command]
	if !exists {
		return fmt.Errorf("unknown command: %s", command)
	}
	for _, status := range validStatuses {
		if status == currentStatus {
			return nil
		}
	}
	return fmt.Errorf("command '%s' is not valid for current status %s", command, currentStatus)
}
