package eventer

import (
	"sync"

	"github.com/labstack/echo/v4"
)

// EventClient interface defines methods for managing event clients and sending events.
type EventClient interface {
	AddClient(userId string) chan string
	RemoveClient(userId string)
	// SendTo sends an event to a specific user identified by userId.
	//
	// It sent to message channel associated with the userId.
	// If the user is not found it returns an error
	SendTo(userId string, event string) error

	// Listen handles incoming requests to listen for events for a specific user.
	//
	// It retrieves the userId from the request context, validates it, and sets up
	// a server-sent events (SSE) stream to send events to the client.
	// If the userId is missing or invalid, it returns a bad request error.
	// It uses the Echo framework's context for handling the request and response.
	Listen(ctx echo.Context) error
}

func NewEventClient() EventClient {
	return &eventManager{
		clients: make(map[string]chan string),
		mutex:   sync.RWMutex{},
		eventID: 0,
	}
}
