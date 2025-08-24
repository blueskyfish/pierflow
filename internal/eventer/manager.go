package eventer

import (
	"errors"
	"fmt"
	"net/http"
	"pierflow/internal/logger"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type eventManager struct {
	clients map[string]chan string
	mutex   sync.RWMutex
	eventID int64
}

func (em *eventManager) AddClient(userId string) chan string {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if _, exists := em.clients[userId]; !exists {
		em.clients[userId] = make(chan string) // blocked channel
		logger.Debugf("[%s] Added new client", userId)
	}
	return em.clients[userId]
}

func (em *eventManager) RemoveClient(userId string) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if ch, exists := em.clients[userId]; exists {
		close(ch)
		delete(em.clients, userId)
		logger.Debugf("[%s] Removed client", userId)
	}
}

func (em *eventManager) SendTo(userId string, event string) error {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	if ch, exists := em.clients[userId]; exists {
		select {
		// blocking send
		case ch <- event:
			return nil
		default:
			return nil
		}
	}
	return errors.New(fmt.Sprintf("user '%s' not found", userId))
}

func (em *eventManager) Listen(ctx echo.Context) error {
	userId := ctx.Param("id")
	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, &map[string]string{
			"message": "user is required",
		})
	}

	ch := em.AddClient(userId)
	// Remove client on exit
	defer em.RemoveClient(userId)

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
		case event, ok := <-ch:
			if !ok {
				logger.Debugf("[%s] Channel closed for user", userId)
				continue
			}

			em.eventID++
			// SSE-Format: id, event type und data
			response := fmt.Sprintf("id: %d\nevent: message\ndata: %s\n\n", em.eventID, event)
			_, err := ctx.Response().Write([]byte(response))
			if err != nil {
				logger.Errorf("[%s] Failed to write event to user: %v", userId, err)
				return err
			}
			ctx.Response().Flush()
			logger.Debugf("[%s] Sent event: %s", userId, event)

		case <-ticker.C:
			// Heartbeat send
			em.eventID++
			response := fmt.Sprintf("id: %d\nevent: heartbeat\ndata: ping\n\n", em.eventID)
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
