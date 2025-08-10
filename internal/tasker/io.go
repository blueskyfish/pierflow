package tasker

import "bytes"

type taskIO struct {
	buf bytes.Buffer
}

func newTaskIO() *taskIO {
	return &taskIO{
		buf: bytes.Buffer{},
	}
}

func (t *taskIO) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t *taskIO) Reset() {
	t.buf.Reset()
}

func (t *taskIO) Read(p []byte) (n int, err error) {
	return t.buf.Read(p)
}

func (t *taskIO) String() string {
	return t.buf.String()
}
