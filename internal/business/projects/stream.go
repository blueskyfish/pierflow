package projects

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"pierflow/internal/logger"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func parseMessage(message string) []Message {
	parts := strings.SplitN(message, "|", 2)
	if len(parts) < 2 {
		return []Message{{
			Status:  "warning",
			Message: fmt.Sprintf("Invalid message format (%s)", message),
			Time:    time.Now(),
		}}
	}

	list := strings.Split(parts[1], "\n")

	var messageList []Message
	for _, msg := range list {
		m := strings.Trim(msg, " \n\t")
		if m == "" {
			continue // Skip empty messages
		}
		messageList = append(messageList, Message{Status: parts[0], Message: m, Time: time.Now()})
	}

	return messageList
}

// receiveMessageAndSent reads messages from the message channel and sends them to the client as a JSON stream.
func receiveMessageAndSent(ctx echo.Context, messageChan chan string, finishFunc func() error) error {
	res := ctx.Response()
	res.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res.Header().Set("Transfer-Encoding", "chunked")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-open")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	encoder := json.NewEncoder(res)
	flusher, ok := res.Writer.(http.Flusher)
	if !ok {
		ctx.Logger().Error("Response writer does not support flushing")
		return errors.New("response writer does not support flushing")
	}

	for {
		message, hasMessage := <-messageChan
		if !hasMessage {
			// Channel closed, finish the response
			if err := finishFunc(); err != nil {
				ctx.Logger().Errorf("Failed to finish response: %s", err.Error())
				return err
			}
			res.WriteHeader(http.StatusOK)
			return nil
		}

		msgList := parseMessage(message)
		for _, msg := range msgList {
			if err := encoder.Encode(&msg); err != nil {
				return err
			}
			logger.Debugf("Sent message to client: %s", msg.Message)
		}
		flusher.Flush() // Ensure the message is sent immediately
	}
}
