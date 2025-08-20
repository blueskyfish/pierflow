package tasker

import (
	"bytes"
	"pierflow/internal/logger"
)

type taskIO struct {
	buf bytes.Buffer
}

func newTaskIO() *taskIO {
	return &taskIO{
		buf: bytes.Buffer{},
	}
}

func (t *taskIO) Write(p []byte) (n int, err error) {
	logger.Debugf("Task writing into buffer: %s", string(p))
	return t.buf.Write(p)
}

func (t *taskIO) Reset() {
	t.buf.Reset()
}

func (t *taskIO) Read(p []byte) (n int, err error) {
	logger.Debug("Task reading from buffer")
	return t.buf.Read(p)
}

func (t *taskIO) String() string {
	return t.buf.String()
}
