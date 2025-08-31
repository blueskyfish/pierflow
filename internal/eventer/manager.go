package eventer

import (
	"errors"
	"fmt"
	"net/http"
	"pierflow/internal/business/utils"
	"pierflow/internal/logger"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type eventServe struct {
	clients map[string]chan ServerSentEvent
	mutex   sync.RWMutex
	eventID int64
}

func (es *eventServe) addClient(userId string) chan ServerSentEvent {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	if _, exists := es.clients[userId]; !exists {
		es.clients[userId] = make(chan ServerSentEvent) // blocked channel
		logger.Debugf("[%s] Added new client", userId)
	}
	return es.clients[userId]
}

func (es *eventServe) removeClient(userId string) {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	if eventChan, exists := es.clients[userId]; exists {
		close(eventChan)
		delete(es.clients, userId)
		logger.Debugf("[%s] Removed client", userId)
	}
}

func (es *eventServe) Send(userId, message, data string) error {
	es.mutex.RLock()
	defer es.mutex.RUnlock()

	if ch, exists := es.clients[userId]; exists {
		select {
		// blocking send
		case ch <- ServerSentEvent{Message: message, Data: data}:
			logger.Debugf("[%s] Sent message to user: %s", userId, message)
			return nil
		default:
			return nil
		}
	}
	return errors.New(fmt.Sprintf("user '%s' not found", userId))
}

func (es *eventServe) Listen(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		userId = ctx.QueryParam("id")
		if userId == "" {
			return ctx.JSON(http.StatusBadRequest, &map[string]string{
				"message": "user is required",
			})
		}
	}

	eventChan := es.addClient(userId)
	// Remove client on exit
	defer es.removeClient(userId)

	// set server-side event Headers
	ctx.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	ctx.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	ctx.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	ctx.Response().WriteHeader(http.StatusOK)

	// TODO make heartbeat interval configurable (default: 30s)
	// start a heartbeat timer every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case event, ok := <-eventChan:
			if !ok {
				logger.Debugf("[%s] Channel closed for user", userId)
				continue
			}

			message := event.Message
			if message == "" {
				message = "message"
			}
			data := event.Data

			es.eventID++
			// SSE-Format: id, event type und data
			response := fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", es.eventID, message, data)
			_, err := ctx.Response().Write([]byte(response))
			if err != nil {
				logger.Errorf("[%s] Failed to write event to user: %v", userId, err)
				return err
			}
			ctx.Response().Flush()
			logger.Debugf("[%s] Sent event: %v", userId, event)

		case <-ticker.C:
			// Heartbeat sends
			es.eventID++
			response := fmt.Sprintf("id: %d\nevent: heartbeat\ndata: ping\n\n", es.eventID)
			_, err := ctx.Response().Write([]byte(response))
			if err != nil {
				logger.Errorf("[%s] Failed to send heartbeat to user: %v", userId, err)
				return err
			}
			ctx.Response().Flush()

		case <-ctx.Request().Context().Done():
			logger.Debugf("[%s] Request is cancel: %v", userId, ctx.Request().Context().Err())
			return nil
		}
	}
}

func (es *eventServe) Messager(userId, message, projectId string, finishFunc func()) Messager {
	if message == "" {
		message = "message"
	}

	eventChan, okay := es.clients[userId]
	if !okay {
		return nil
	}

	bodyChan := make(chan MessageBody) // blocked channel
	go es.broadcast(bodyChan, message, eventChan, finishFunc)

	return &messager{
		status:    StatusDebug,
		projectId: projectId,
		channel:   bodyChan,
		TimeFunc:  defaultTimeFunc,
	}
}

func (es *eventServe) broadcast(channel chan MessageBody, message string, eventChan chan ServerSentEvent, finishFunc func()) {
	for {
		msg, ok := <-channel
		if !ok {
			logger.Debugf("[%s] Channel closed for user", message)
			if finishFunc != nil {
				finishFunc()
			}
			return
		}

		data, err := utils.Stringify(msg)
		if err != nil {
			logger.Errorf("Failed to stringify message: %v", err)
			continue
		}

		event := ServerSentEvent{
			Message: message,
			Data:    data,
		}
		eventChan <- event
	}
}
