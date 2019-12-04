package txt2byte

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

func (c coder) Rx(data interface{}, ctx *handler.Context) (err error) {

	if d, ok := data.(string); ok {
		ctx.TxWrite([]byte(d))
	}
	return
}

func (c *coder) Tx(data interface{}, ctx *handler.Context) (err error) {

	if d, ok := data.([]byte); ok {
		ctx.RxWrite(string(d))
	}
	return
}
