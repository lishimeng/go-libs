package pipeline

import (
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
	"io"
)

type Pipeline struct {
	tx     *handler.TxContainer
	rx     *handler.RxContainer
	writer io.Writer
}

type MessageContext struct {
	payload interface{}
	meta    *handler.Meta
}

func newPipeline(c handler.Container) (p *Pipeline) {

	enc := &handler.TxContainer{
		Container: c,
	}
	dec := handler.RxContainer{
		Container: c,
	}
	p = &Pipeline{
		tx: enc,
		rx: &dec,
	}
	return
}
