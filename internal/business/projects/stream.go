package projects

import (
	"fmt"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"
	"strings"
	"time"
)

func parseMessage(projectId, action, message string) []EventMessageResponse {
	parts := strings.SplitN(message, "|", 2)
	if len(parts) < 2 {
		return []EventMessageResponse{{
			Action:    action,
			ProjectId: projectId,
			Status:    "warning",
			Message:   fmt.Sprintf("Invalid message format (%s)", message),
			Time:      time.Now(),
		}}
	}

	list := strings.Split(parts[1], "\n")

	var messageList []EventMessageResponse
	for _, msg := range list {
		m := strings.Trim(msg, " \n\t")
		if m == "" {
			continue // Skip empty messages
		}
		messageList = append(messageList, EventMessageResponse{
			Action:    action,
			ProjectId: projectId,
			Status:    parts[0],
			Message:   m,
			Time:      time.Now(),
		})
	}

	return messageList
}

type receiveOptions struct {
	userId      string       // user unique id
	projectId   string       // project unique id
	action      string       // the action name of the command
	messageChan chan string  // the message channel to read messages from
	finishFunc  func() error // function to call when the channel is closed
}

func buildReceiveOptions(userId, projectId, action string, messageChan chan string, finishFunc func() error) receiveOptions {
	return receiveOptions{
		userId:      userId,
		projectId:   projectId,
		action:      action,
		messageChan: messageChan,
		finishFunc:  finishFunc,
	}
}

// receiveMessageAndSent reads messages from the message channel and sends then to the user via server side events.
//
// When the channel is closed, it calls the finishFunc to perform any final actions (like updating project status) and ends the response.
func (pm *ProjectManager) receiveMessageAndSent(options receiveOptions) error {
	logger.Debugf("Build project: Listening for messages for user %s", options.userId)
	for {
		message, hasMessage := <-options.messageChan
		if !hasMessage {
			// Channel closed, finish the response
			if err := options.finishFunc(); err != nil {
				logger.Errorf("Failed to finish response: %s", err.Error())
				return err
			}
			logger.Debugf("Message channel closed for user '%s'", options.userId)
			if message == "" {
				return nil
			}
		}

		msgList := parseMessage(options.projectId, options.action, message)
		for _, msg := range msgList {
			sendMessage, err := utils.Stringify(msg)
			if err != nil {
				logger.Errorf("Failed to serialize message: %s", err.Error())
				continue
			}
			err = pm.eventClient.SendTo(options.userId, sendMessage)
			if err != nil {
				logger.Errorf("Failed to send message to user '%s': %v", options.userId, err)
				continue
			}
		}
	}
}
