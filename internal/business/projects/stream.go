package projects

import (
	"pierflow/internal/business/utils"
	"pierflow/internal/eventer"
	"pierflow/internal/logger"
)

func (pm *ProjectManager) listenEventMessager(userId, projectId, action string, messager eventer.Messager, finishFunc func() error) error {
	if finishFunc == nil {
		// add a default finish function to avoid nil pointer dereference
		finishFunc = func() error { return nil }
	}
	foundError := false
	for {
		hasMessage, msg := messager.Receive()
		if !hasMessage {
			// Channel closed, finish the response
			logger.Debugf("Messager channel closed for user '%s'", userId)
			if msg == nil {
				if foundError {
					return nil
				} else if err := finishFunc(); err != nil {
					// call finish function
					logger.Errorf("Failed to finish response: %s", err.Error())
					return err
				}

				// leave if the channel is closed and no message
				return nil
			}
		}
		foundError = foundError || (msg != nil && msg.Status == eventer.StatusError)
		sendMessage, err := utils.Stringify(toEventMessageResponse(projectId, action, msg))
		if err != nil {
			logger.Errorf("Failed to serialize message: %s", err.Error())
			continue
		}
		err = pm.eventClient.SendTo(userId, "", sendMessage)
		if err != nil {
			logger.Errorf("Failed to send message to user '%s': %v", userId, err)
			continue
		}
	}
}
