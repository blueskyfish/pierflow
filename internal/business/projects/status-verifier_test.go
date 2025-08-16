package projects

import "testing"

func TestStatusVerifier_VerifyStatus(t *testing.T) {

	tests := []struct {
		command       ProjectCommand
		currentStatus ProjectStatus
		expectedError bool
	}{
		{CommandCreateProject, StatusCreated, false},
		{CommandCloneRepository, StatusCreated, false},
		{CommandCheckoutRepository, StatusCloned, false},
		{CommandBuildProject, StatusCheckedOut, false},
		{CommandPullRepository, StatusCloned, false},
		{CommandStartProject, StatusBuilt, false},
		{CommandStopProject, StatusRun, false},
		{CommandBuildProject, StatusStopped, false},
		{CommandCheckoutRepository, StatusStopped, true},
		{CommandDeleteProject, StatusStopped, false},
		{CommandDeleteProject, StatusCloned, false},
		{CommandCreateProject, StatusCloned, false}, // Create project command does not have a status
		{CommandCloneRepository, StatusCheckedOut, true},
		{CommandCheckoutRepository, StatusBuilt, true},
		{CommandBuildProject, StatusRun, true},
		{CommandPullRepository, StatusRun, true},
		{CommandStartProject, StatusCloned, true},
		{CommandStopProject, StatusCloned, true},
		{CommandDeleteProject, StatusRun, true},
		{ProjectCommand("UnknownCommand"), StatusCreated, true}, // Testing an unknown command
	}

	for _, tt := range tests {
		t.Run(tt.command.String(), func(t *testing.T) {
			err := verifier.VerifyStatus(tt.command, tt.currentStatus)
			if (err != nil) != tt.expectedError {
				t.Errorf("VerifyStatus() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
