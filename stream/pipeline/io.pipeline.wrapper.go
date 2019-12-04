package pipeline

import (
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
	"github.com/lishimeng/go-libs/stream/reactor"
	"io"
)

func New(rwc io.ReadWriteCloser, container handler.Container) (pipe *Pipeline, r *reactor.Stream) {

	pipe = newPipeline(container)
	r = reactor.New(rwc)
	r.DataListener(pipe.OnRx) // pipeline <- reactor
	pipe.writer = r.Writer    // reactor <- pipeline

	return
}
