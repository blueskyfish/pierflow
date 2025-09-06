package eventer

import (
	"sync"

	"github.com/labstack/echo/v4"
)

// ServerSentEvent represents the structure of an event message.
//
// if the data field `EventType` is empty, the event is considered the default event without a `message? type
type ServerSentEvent struct {
	EventType string // The message type e.g. "message", "update", "delete", etc.
	Data      string // The actual data payload, typically a JSON string
}

// EventServe interface defines methods for managing event clients and sending events.
type EventServe interface {
	// Send sends an event to a specific user identified by userId.
	//
	// It sent to a message channel associated with the userId.
	// If the user is not found, it returns an error
	//
	// The Send method is called in the producer goroutine to send events to the client.
	// It is a blocking call, meaning it will wait until the event is sent to the channel.
	// If the userId does not exist, it returns an error indicating that the user was not found.
	Send(userId, message, data string) error

	// Listen handles incoming requests to listen for events for a specific user.
	//
	// It retrieves the userId from the request context, validates it, and sets up
	// a server-sent events (SSE) stream to send events to the client.
	// If the userId is missing or invalid, it returns a bad request error.
	// It uses the Echo framework's context for handling the request and response.
	//
	// The Listen method is called as a consumer function to receive events from the channel and send them to the client.
	Listen(ctx echo.Context) error

	// WithMessage creates a new Messager instance for sending messages (MessageBody) to a specific user.
	//
	// The project id is for the consumer to filter messages. The status is the default status for messages sent by
	// this Messager. It is allowed to be empty, in which case "debug" is used as the default status.
	//
	// Before a specific command is executed, a Messager should be created to capture all messages
	// related to that command. The Messager should be closed after the command is completed to avoid memory leaks.
	// The finishFunc is a function that will be called when the Messager is closed.
	//
	// The Messager method is called in the producer goroutine to create a new message channel for sending messages.
	// It returns a Messager instance that can be used to send messages to the client.
	WithMessage(eventType, userId, projectId string, finishFunc func(data interface{})) Messager
}

func NewEventServe() EventServe {
	return &eventServe{
		clients: make(map[string]chan ServerSentEvent),
		mutex:   sync.RWMutex{},
		eventID: 0,
	}
}
