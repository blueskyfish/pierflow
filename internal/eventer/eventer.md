
# Pierflow Eventer

## Usage

```go
messager := pm.eventer.WithMessage(
    "command.build",
    eventer.StatusDebug,
)

err := messager.Send(eventer.StatusSuccess, projectId, toHead(head))
if err != nil {
	...
}
```

## Entity

Das Event wird in ein Server-Side Event gepackt mit einem Message gepackt.

### Message

| Name      | Type      | Json      |
|-----------|-----------|-----------|
| Status    | string    | status    |
| Data      | string    | data      |
| ProjectId | string    | projectId |
| Time      | time.Time | time      |


### ServerSentEvent

| Name    | Type    | Json  |
|---------|---------|-------|
| Message | string  | event |
| Data    | Message | data  |


## Handler

### Connect(ctx echo.Context) err

Handle the server side events initialized by the browser with a unique id.

### WithEvent(event, userId, projectId string) Messager

Create a messager instance to push events from the server to the client

## Messager
