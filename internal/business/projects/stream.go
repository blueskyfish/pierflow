package projects

import (
	"fmt"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"
	"strings"
	"time"
)

func parseMessage(projectId, action, message string) []Message {
	parts := strings.SplitN(message, "|", 2)
	if len(parts) < 2 {
		return []Message{{
			Action:    action,
			ProjectId: projectId,
			Status:    "warning",
			Message:   fmt.Sprintf("Invalid message format (%s)", message),
			Time:      time.Now(),
		}}
	}

	list := strings.Split(parts[1], "\n")

	var messageList []Message
	for _, msg := range list {
		m := strings.Trim(msg, " \n\t")
		if m == "" {
			continue // Skip empty messages
		}
		messageList = append(messageList, Message{
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
	userId      string
	projectId   string
	action      string
	messageChan chan string
	finishFunc  func() error
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
// When the channel is closed, it calls the finishFunc to perform any final actions (like updating project status) and ends the response.
func receiveMessageAndSent(options receiveOptions) error {
	logger.Debugf("Build project: Listening for messages for user %s", options.userId)
	for {
		message, hasMessage := <-options.messageChan
		if !hasMessage {
			// Channel closed, finish the response
			if err := options.finishFunc(); err != nil {
				logger.Errorf("Failed to finish response: %s", err.Error())
				return err
			}
			logger.Debugf("Message channel closed for user %s", options.userId)
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
			err = connManager.SendTo(options.userId, sendMessage)
			if err != nil {
				logger.Errorf("Failed to send message to user %s: %s", options.userId, err.Error())
				continue
			}
		}
	}
}
