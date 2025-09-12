package docker

import (
	"context"
	"errors"
	"pierflow/internal/logger"

	"github.com/moby/moby/api/types/events"
	"github.com/moby/moby/client"
)

type ComposeClient interface {
	Listen(ctx context.Context, eventChan chan<- ComposeEvent) error
}

type composeClient struct {
	actionFilters []events.Action
}

// NewComposeClient creates a new ComposeClient with optional action filters.
//
// If actionFilters is empty, all compose events are listened to.
func NewComposeClient(actionFilters []events.Action) ComposeClient {
	return &composeClient{
		actionFilters: actionFilters,
	}
}

func (c *composeClient) Listen(ctx context.Context, eventChan chan<- ComposeEvent) error {
	if ctx == nil {
		return errors.New("context is required")
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	go c.runListen(ctx, cli, eventChan)

	return nil
}

func (c *composeClient) runListen(ctx context.Context, cli *client.Client, eventChan chan<- ComposeEvent) {
	defer c.Close(cli)

	msgChan, errChan := cli.Events(ctx, client.EventsListOptions{})

	for {
		select {
		case msg := <-msgChan:
			if !isComposeMessage(msg, c.actionFilters) {
				continue
			}
			eventChan <- toComposeEvent(msg)
		case err := <-errChan:
			logger.Errorf("error listening to docker events: %v", err)
		case <-ctx.Done():
			return
		}
	}
}

func (c *composeClient) Close(cli *client.Client) {
	err := cli.Close()
	if err != nil {
		logger.Errorf("close compose client with error: %v", err)
	}
}
