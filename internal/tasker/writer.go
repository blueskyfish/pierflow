package tasker

import "fmt"

const (
	MessageError = -1
	MessageOK    = 0
	MessageEnd   = 1
)

type TaskWriter struct {
	prefix      string
	messageChan chan string
}

func newTaskWriter(prefix string, messageChan chan string) *TaskWriter {
	return &TaskWriter{
		prefix:      prefix,
		messageChan: messageChan,
	}
}

func (tw *TaskWriter) Write(b []byte) (num int, err error) {
	message := string(b)
	if len(message) == 0 {
		return 0, nil
	}
	if message == "\n" {
		return len(b), nil
	}

	tw.messageChan <- toMessage(tw.prefix, message)
	return len(b), nil
}

func toMessage(prefix, message string) string {
	return fmt.Sprintf("%s|%s", prefix, message)
}

func toOkay(message string) string {
	return toMessage("okay", message)
}

func toOkayf(format string, params ...any) string {
	return toOkay(fmt.Sprintf(format, params...))
}

func toError(message string) string {
	return toMessage("error", message)
}

func toErrorf(format string, params ...any) string {
	return toError(fmt.Sprintf(format, params...))
}
