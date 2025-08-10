package gitter

import (
	"bytes"
	"io"
)

type Progressor struct {
	buf bytes.Buffer
}

func newProgressor() *Progressor {
	return &Progressor{
		buf: bytes.Buffer{},
	}
}

func (p *Progressor) Write(data []byte) (int, error) {
	return p.buf.Write(data)
}

func (p *Progressor) String() string {
	return p.buf.String()
}

func (p *Progressor) Reset() {
	p.buf.Reset()
}

func (p *Progressor) Reader() io.Reader {
	return &p.buf
}
