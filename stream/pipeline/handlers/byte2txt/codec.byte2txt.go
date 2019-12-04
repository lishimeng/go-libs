package byte2txt

import (
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
)

type coder struct {
}

func New() (h handler.Handler) {

	var t handler.Handler = &coder{}
	h = t
	return h
}

func (c coder) Tx(input interface{}, ctx *handler.Context) (err error) {

	if in, ok := input.(string); ok {
		ctx.TxWrite([]byte(in))
	}
	return
}

func (c *coder) Rx(input interface{}, ctx *handler.Context) (err error) {

	if in, ok := input.([]byte); ok {
		ctx.RxWrite(string(in))
	}
	return
}
