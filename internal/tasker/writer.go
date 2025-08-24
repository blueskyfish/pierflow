package tasker

import "fmt"

type TaskWriter struct {
	prefix      string
	messageChan chan string
}

// newTaskWriter creates a new TaskWriter with the given prefix and message channel.
//
// The prefix is used to categorize the messages (e.g., "okay", "error").
// The messageChan is a channel where the formatted messages will be sent `e.g. "okay|Task completed successfully"`.
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
