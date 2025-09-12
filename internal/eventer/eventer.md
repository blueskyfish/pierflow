
# Pierflow Eventer

## Messaging System and Status

The messager inside contains a messaging system that allows sending messages with different statuses. The statuses are
defined as constants.

* StatusDebug
* StatusInfo
* StatusWarn
* StatusError
* **StatusSuccess** this is special as they are only sent once the command is finished and it send the result.

## Usage

```go
package projects

import (
	"github.com/blueskyfish/pierflow/internal/business/projects"
	"github.com/blueskyfish/pierflow/internal/business/utils"
	"github.com/blueskyfish/pierflow/internal/eventer"

	"github.com/labstack/echo/v4"
)

func (pm *ProjectManager) ACommand(ctx echo.Context) {
	userId := utils.HeaderUser(ctx)
	projectId := ctx.Param("id")
	var messager = pm.eventServe.WithMessage(
		projects.CommandBuildProject.Message(),
		userId,
		projectId,
		func(data interface{}) {
			// do something with data
		},
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