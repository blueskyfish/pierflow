package docker

import "github.com/moby/moby/api/types/events"

// isComposeMessage checks if the event is related to a Docker Compose managed container
// by looking for the "com.docker.compose.project" attribute.
func isComposeMessage(event events.Message, actionFilters []events.Action) bool {
	if event.Type != "container" {
		return false
	}
	_, ok := event.Actor.Attributes["com.docker.compose.project"]
	if ok {
		_, ok = event.Actor.Attributes["com.docker.compose.project.working_dir"]
	}
	if ok {
		_, ok = event.Actor.Attributes["com.docker.compose.service"]
	}
	if ok {
		_, ok = event.Actor.Attributes["com.docker.compose.container-number"]
	}
	if ok {
		if len(actionFilters) == 0 {
			return true
		}
		for _, filter := range actionFilters {
			if event.Action == filter {
				return true
			}
		}
		return false
	}
	return ok
}

// toComposeEvent transforms a Docker event.Message into a ComposeEvent,
// extracting relevant attributes for easier access.
func toComposeEvent(ev events.Message) ComposeEvent {
	return ComposeEvent{
		Action:     ev.Action,
		ID:         ev.Actor.ID,
		Project:    ev.Actor.Attributes["com.docker.compose.project"],
		WorkingDir: ev.Actor.Attributes["com.docker.compose.project.working_dir"],
		Image:      ev.Actor.Attributes["image"],
		Name:       ev.Actor.Attributes["name"],
		Service:    ev.Actor.Attributes["com.docker.compose.service"],
		Container:  ev.Actor.Attributes["com.docker.compose.container-number"],
		Time:       ev.Time, // MAYBE use TimeNano for more precision
	}
}
