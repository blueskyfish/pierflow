package serve

import (
	"strings"

	"github.com/moby/moby/api/types/events"
)

func parseActions(actions string) []events.Action {
	parts := strings.Split(actions, ",")
	var result []events.Action
	for _, part := range parts {
		action := strings.TrimSpace(part)
		switch strings.ToLower(action) {
		case "create", "start", "restart", "stop",
			"checkpoint", "pause", "unpause", "attach",
			"detach", "resize", "update", "rename",
			"kill", "die", "oom", "destroy",
			"remove", "commit", "top", "copy",
			"archive-path", "extract-to-dir",
			"export", "import", "save", "load", "tag",
			"untag", "push", "pull", "prune",
			"delete", "enable", "disable", "connect",
			"disconnect", "reload", "mount", "unmount":

		default:
			continue
		}

		result = append(result, events.Action(strings.ToLower(action)))
	}

	if len(result) == 0 {
		return []events.Action{events.ActionStart, events.ActionRestart, events.ActionStop}
	}
	return result
}
