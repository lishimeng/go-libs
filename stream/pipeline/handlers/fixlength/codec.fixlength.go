package fixlength

import (
	"bytes"
	"github.com/lishimeng/go-libs/log"
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
)

type coder struct {
	handler.Context
	length    int
	decodeBuf *bytes.Buffer
	encodeBuf *bytes.Buffer
}

func New(length int) handler.Handler {

	coder := &coder{
		Context:   *handler.NewContext(8, 8),
		length:    length,
		decodeBuf: new(bytes.Buffer),
		encodeBuf: new(bytes.Buffer),
	}
	var c handler.Handler = coder
	return c
}

func (c *coder) Rx(input interface{}) (err error) {

	log.Fine("byte raw decode")
	if in, ok := input.([]byte); ok {
		c.decodeBuf.Write(in)
	}

	for c.decodeBuf.Len() >= c.length {
		log.Fine("build a frame")

		p := make([]byte, c.length)
		_, err = c.decodeBuf.Read(p)
		if err != nil {
			return
		}
		c.RxWrite(p)
	}
	return
}

func (c coder) Tx(input interface{}) (err error) {

	if in, ok := input.([]byte); ok {
		c.encodeBuf.Write(in)
	}
	for c.encodeBuf.Len() >= c.length {

		var p = make([]byte, c.length)
		_, err = c.encodeBuf.Read(p)
		if err != nil {
			return
		}
		c.TxWrite(p)
	}
	return
}
