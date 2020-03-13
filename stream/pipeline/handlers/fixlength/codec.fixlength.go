package fixlength

import (
	"bytes"
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
)

type coder struct {
	length    int
	decodeBuf *bytes.Buffer
	encodeBuf *bytes.Buffer
}

func New(length int) handler.Handler {

	coder := &coder{
		length:    length,
		decodeBuf: new(bytes.Buffer),
		encodeBuf: new(bytes.Buffer),
	}
	var c handler.Handler = coder
	return c
}

func (c *coder) Rx(input interface{}, ctx *handler.Context) (err error) {

	if in, ok := input.([]byte); ok {
		c.decodeBuf.Write(in)
	}

	for c.decodeBuf.Len() >= c.length {

		p := make([]byte, c.length)
		_, err = c.decodeBuf.Read(p)
		if err != nil {
			return
		}
		ctx.RxWrite(p)
	}
	return
}

func (c coder) Tx(input interface{}, ctx *handler.Context) (err error) {

	if in, ok := input.([]byte); ok {
		c.encodeBuf.Write(in)
	}
	for c.encodeBuf.Len() >= c.length {

		var p = make([]byte, c.length)
		_, err = c.encodeBuf.Read(p)
		if err != nil {
			return
		}
		ctx.TxWrite(p)
	}
	return
}
