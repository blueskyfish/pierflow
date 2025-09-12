package projects

import (
	"context"
	"github.com/blueskyfish/pierflow/internal/docker"
	"github.com/blueskyfish/pierflow/internal/logger"
)

// ListenForComposeEvents starts listening for Docker Compose events and handles them.
//
// It spawns a goroutine to process incoming events.
func (pm *ProjectManager) ListenForComposeEvents() error {
	ctx := context.Background()

	composeEventChan := make(chan docker.ComposeEvent)
	err := pm.composeClient.Listen(ctx, composeEventChan)
	if err != nil {
		return err
	}

	go pm.runListenForComposeEvents(composeEventChan)
	return nil
}

func (pm *ProjectManager) runListenForComposeEvents(eventChan <-chan docker.ComposeEvent) {
	for event := range eventChan {
		logger.Infof("Received compose event: %+v", event)
		// which project does this event belong to?
	}
}
