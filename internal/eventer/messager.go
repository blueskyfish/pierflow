package eventer

import (
	"encoding/json"
	"errors"
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

type Message struct {
	Status  string    `json:"status"`
	Message string    `json:"message"` // can be any type, using interface{} for flexibility
	Time    time.Time `json:"time"`    // The UTC timestamp of the message
}

type TimeFunc func() time.Time

type Messager interface {
	io.Writer
	io.Closer

	// Send sends a message with the given status and message content.
	//
	// Send is called in the producer goroutine.
	// The message parameter can be of any type; if it's not a string, it will be marshaled to JSON.
	Send(status string, message interface{}) error

	// Receive receives a message from the channel.
	//
	// Receive is called in the consumer goroutine.
	// It returns a boolean indicating if the channel is open and a pointer to the Message.
	Receive() (bool, *Message)

	// Closing closes the message channel.
	//
	// It is called in the producer goroutine to signal that no more messages will be sent.
	Closing()
}

// NewMessager creates a new Messager instance with the provided TimeFunc.
//
// The status parameter is the default status for messages; it can be overridden in the Send method.
// The TimeFunc is used to generate timestamps for messages; if nil, the current UTC time is used.
func NewMessager(status string, timeFunc TimeFunc) Messager {
	if timeFunc == nil {
		timeFunc = defaultTimeFunc
	}
	return &messager{
		status:   status,
		channel:  make(chan Message), // Blocked channel
		TimeFunc: timeFunc,
	}
}

// messager implements the Messager interface using a channel for communication.
type messager struct {
	status   string // The default status for messages
	channel  chan Message
	TimeFunc TimeFunc
}

func (m *messager) Send(status string, message interface{}) error {
	if _, ok := message.(string); !ok {
		byteValue, err := json.Marshal(message)
		if err != nil {
			return errors.New("failed to convert message instance to string")
		}
		message = string(byteValue)
	}

	msg := Message{
		Status:  status,
		Message: message.(string),
		Time:    m.TimeFunc(),
	}
	m.channel <- msg
	return nil
}

func (m *messager) Receive() (bool, *Message) {
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
	if data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	if len(data) == 0 {
		return len(data), nil
	}

	messages := strings.Split(string(data), "\n")
	for _, message := range messages {
		err = m.Send(m.status, message)
		if err != nil {
			return 0, err
		}
	}

	return len(data), nil
}

func (m *messager) Close() error {
	close(m.channel)
	return nil
}

func (m *messager) Closing() {
	_ = m.Close()
}

func defaultTimeFunc() time.Time {
	return time.Now().UTC()
}
