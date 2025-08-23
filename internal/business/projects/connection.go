package projects

import (
	"errors"
	"fmt"
	"net/http"
	"pierflow/internal/logger"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type ConnectManager struct {
	clients map[string]chan string
	mutex   sync.RWMutex
}

var connManager = &ConnectManager{
	clients: make(map[string]chan string),
	mutex:   sync.RWMutex{},
}

func (cm *ConnectManager) AddClient(userId string) chan string {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if _, exists := cm.clients[userId]; !exists {
		cm.clients[userId] = make(chan string) // blocked channel
		logger.Debugf("Added new client for user %s", userId)
	}
	return cm.clients[userId]
}

func (cm *ConnectManager) RemoveClient(userId string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if ch, exists := cm.clients[userId]; exists {
		close(ch)
		delete(cm.clients, userId)
		logger.Debugf("Removed client for user %s", userId)
	}
}

func (cm *ConnectManager) SendTo(userId string, event string) error {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	if ch, exists := cm.clients[userId]; exists {
		select {
		case ch <- event:
			return nil
		default:
			return nil // Channel is full, drop the event
		}
	}
	return errors.New(fmt.Sprintf("user '%s' not found", userId))
}

func (cm *ConnectManager) Listen(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("user is required"))
	}

	ch := cm.AddClient(userId)
	defer cm.RemoveClient(userId)

	// set SSE-Header
	ctx.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	ctx.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	ctx.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	ctx.Response().WriteHeader(http.StatusOK)

	// start a heartbeat timer
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Event-ID für SSE
	eventID := 0

	for {
		select {
		case event, ok := <-ch:
			if !ok {
				logger.Debugf("Channel closed for user %s", userId)
				continue
				// return nil // Channel closed
			}

			eventID++
			// SSE-Format: id, event type und data
			response := fmt.Sprintf("id: %d\nevent: message\ndata: %s\n\n", eventID, event)
			_, err := ctx.Response().Write([]byte(response))
			if err != nil {
				logger.Errorf("Failed to write event to user %s: %v", userId, err)
				return err
			}
			ctx.Response().Flush()
			logger.Debugf("Sent event to user '%s': %s", userId, event)

		case <-ticker.C:
			// Heartbeat send
			eventID++
			response := fmt.Sprintf("id: %d\nevent: heartbeat\ndata: ping\n\n", eventID)
			_, err := ctx.Response().Write([]byte(response))
			if err != nil {
				logger.Errorf("Failed to send heartbeat to user '%s': %v", userId, err)
				return err
			}
			ctx.Response().Flush()

		case <-ctx.Request().Context().Done():
			logger.Debugf("Request is cancel for user '%s': %v", userId, ctx.Request().Context().Err())
			return nil
		}
	}
}

func (pm *ProjectManager) UserConnect(ctx echo.Context) error {
	err := connManager.Listen(ctx)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusNoContent, "")
}

func (pm *ProjectManager) SendPing(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, toErrorResponse("user is required"))
	}
	err := connManager.SendTo(userId, "ping")
	if err != nil {
		return ctx.JSON(http.StatusNotFound, toErrorResponse(err.Error()))
	}
	return ctx.String(http.StatusNoContent, "")
}
