package pipeline

import (
	"fmt"
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
)

func (p *Pipeline) OnRx(data []byte) {

	var message = MessageContext{
		payload: data,
		meta:    nil,
	}
	err := p.handleRx(message)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Pipeline) handleRx(message MessageContext) (err error) {

	var nextHandler handler.Handler
	var nextMeta handler.Meta
	var ctx *handler.Context
	if message.meta == nil {
		nextHandler, nextMeta, ctx, err = p.rx.FirstHandler(handler.SortASC)
	} else if p.rx.HasNext(*message.meta, handler.SortASC) {
		nextHandler, nextMeta, ctx, err = p.rx.NextHandler(*message.meta)
	} else {
		return
	}

	if err != nil {
		return
	}
	err = nextHandler.Rx(message.payload, ctx)
	if err != nil {
		return
	}
	var n = ctx.RxLen()
	for i := 0; i < n; i++ {
		var r interface{}
		r, err = ctx.RxRead()
		if err == nil {
			ctx := MessageContext{
				payload: r,
				meta:    &nextMeta,
			}
			err = p.handleRx(ctx)
		}
		if err != nil {
			break
		}
	}
	return
}

func (p *Pipeline) Tx(data interface{}) (err error) {
	var message = MessageContext{
		payload: data,
		meta:    nil,
	}
	err = p.handleTx(message)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (p *Pipeline) handleTx(message MessageContext) (err error) {

	var nextHandler handler.Handler
	var nextMeta handler.Meta
	var ctx *handler.Context
	if message.meta == nil {
		nextHandler, nextMeta, ctx, err = p.tx.FirstHandler(handler.SortDESC)
	} else {
		if p.tx.HasNext(*message.meta, handler.SortDESC) {
			nextHandler, nextMeta, ctx, err = p.tx.NextHandler(*message.meta)
		} else {
			data, isBytes := message.payload.([]byte)
			if isBytes && p.writer != nil {
				_, err = p.writer.Write(data)
			}
			return
		}
	}

	if err != nil {
		return
	}
	err = nextHandler.Tx(message.payload, ctx)

	if err != nil {
		return
	}
	var n = ctx.TxLen()
	for i := 0; i < n; i++ {
		var r interface{}
		r, err = ctx.TxRead()
		if err == nil {
			ctx := MessageContext{
				payload: r,
				meta:    &nextMeta,
			}
			err = p.handleTx(ctx)
		}
		if err != nil {
			break
		}
	}
	return
}
