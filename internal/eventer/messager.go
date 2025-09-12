package eventer

import (
	"github.com/blueskyfish/pierflow/internal/business/utils"
	"github.com/blueskyfish/pierflow/internal/logger"
	"io"
	"strings"
	"time"
)

const (
	StatusInfo    = "info"
	StatusError   = "error"
	StatusDebug   = "debug"
	StatusSuccess = "success"
)

type MessageBody struct {
	Status    string    `json:"status"`
	ProjectId string    `json:"id"`
	Message   string    `json:"message"` // can be any type, using interface{} for flexibility
	Time      time.Time `json:"time"`    // The UTC timestamp of the message
}

type TimeFunc func() time.Time

// Messager send MessageBody to the channel
type Messager interface {
	io.Writer

	// Send sends a message with the given status and message content.
	//
	// Send is called in the producer goroutine.
	// The message parameter can be of any type; if it's not a string, it will be marshaled to JSON.
	Send(status string, message interface{})

	// Receive receives a message from the channel.
	//
	// Receive is called in the consumer goroutine.
	// It returns a boolean indicating if the channel is open and a pointer to the MessageBody.
	Receive() (bool, *MessageBody)

	// Close closes the message channel.
	//
	// It is called in the producer goroutine to signal that no more messages will be sent.
	Close()
}

// messager implements the Messager interface using a channel for communication.
type messager struct {
	status    string           // The default status for messages
	projectId string           // The project ID associated with the messages
	channel   chan MessageBody // The channel for sending messages
	TimeFunc  TimeFunc         // Function to get the current time, defaults to UTC now
}

func (m *messager) Send(status string, message interface{}) {
	if _, ok := message.(string); !ok {
		// Convert non-string message to JSON string
		text, err := utils.Stringify(message)
		if err != nil {
			logger.Warnf("failed to marshal message: %v", err)
			return
		}
		message = text
	}

	msg := MessageBody{
		Status:    status,
		ProjectId: m.projectId,
		Message:   message.(string),
		Time:      m.TimeFunc(),
	}
	m.channel <- msg
}

func (m *messager) Receive() (bool, *MessageBody) {
	msg, ok := <-m.channel
	if !ok {
		return false, nil
	}
	return true, &msg
}

func (m *messager) Write(data []byte) (n int, err error) {
	if len(data) == 0 {
		return 0, nil
	}

	messages := strings.Split(string(data), "\n")
	for _, message := range messages {
		if message == "" {
			continue
		}
		m.Send(m.status, message)
	}

	return len(data), nil
}

func (m *messager) Close() {
	close(m.channel)
}

func defaultTimeFunc() time.Time {
	return time.Now().UTC()
}
