package eventer

import (
	"encoding/json"
	"sync"

	"github.com/labstack/echo/v4"
)

// EventData represents the structure of an event message.
//
// if the data field `Event` is empty, the event is considered default event without message type
type EventData struct {
	Event string
	Data  string
}

// EventClient interface defines methods for managing event clients and sending events.
type EventClient interface {
	AddClient(userId string) chan EventData
	RemoveClient(userId string)

	// SendTo sends an event to a specific user identified by userId.
	//
	// It sent to message channel associated with the userId.
	// If the user is not found it returns an error
	//
	// The SendTo method is called in the producer goroutine to send events to the client.
	// It is a blocking call, meaning it will wait until the event is sent to the channel.
	// If the userId does not exist, it returns an error indicating that the user was not found.
	SendTo(userId, event, data string) error

	// Listen handles incoming requests to listen for events for a specific user.
	//
	// It retrieves the userId from the request context, validates it, and sets up
	// a server-sent events (SSE) stream to send events to the client.
	// If the userId is missing or invalid, it returns a bad request error.
	// It uses the Echo framework's context for handling the request and response.
	//
	// The Listen method is called as consumer function to receive events from the channel and send them to the client.
	Listen(ctx echo.Context) error
}

func NewEventClient() EventClient {
	return &eventManager{
		clients: make(map[string]chan EventData),
		mutex:   sync.RWMutex{},
		eventID: 0,
	}
}

func toChannel(message string, data interface{}) (*EventData, error) {
	if message == "" {
		message = "message"
	}

	dataStr, ok := data.(string)
	if !ok {
		value, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return &EventData{Event: message, Data: string(value)}, nil
	}
	return &EventData{Event: message, Data: dataStr}, nil
}
