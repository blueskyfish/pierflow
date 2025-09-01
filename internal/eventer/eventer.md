
# Pierflow Eventer

## Usage

```go
package projects

import (
	"pierflow/internal/business/projects"
	"pierflow/internal/eventer"
)

func (pm *ProjectManager) ACommand() {
	var messager = pm.eventServe.WithMessage(
		projects.CommandBuildProject.Message(),
		userId,
		projectId,
	)

	err := messager.Send(eventer.StatusSuccess, projectId, toHead(head))
	if err != nil {
		// ...
	}
}
```

## Entity

Das ServerSentEvent entity for sending events to the client. It contains a MessageBody as Data. The MessageBody is
a JSON serializable struct that contains the status, data, projectId and time.

### ServerSentEvent

| Name    | Type        | Json  |
|---------|-------------|-------|
| Message | string      | event |
| Data    | MessageBody | data  |

### MessageBody

| Name      | Type      | Json      |
|-----------|-----------|-----------|
| Status    | string    | status    |
| Message   | string    | data      |
| ProjectId | string    | projectId |
| Time      | time.Time | time      |

> The `Message` field is a string to allow any JSON serializable struct to be sent as data.

## Status

| Name          | Value     |
|---------------|-----------|
| StatusDebug   | "debug"   |
| StatusInfo    | "info"    |
| StatusWarn    | "warn"    |
| StatusError   | "error"   |
| StatusSuccess | "success" |

## EventType

Every command has an event type. The command type has a prefix of "message-" and extends with the command name.

> e.g. "message-build-project"