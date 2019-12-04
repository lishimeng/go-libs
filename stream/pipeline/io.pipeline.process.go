package pipeline

import (
	"encoding/hex"
	"fmt"
	"github.com/lishimeng/go-libs/log"
	"github.com/lishimeng/go-libs/stream/pipeline/handler"
)

func (p *Pipeline) OnRx(data []byte) {

	log.Fine("OnRx:%s", hex.EncodeToString(data))
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
	if message.meta == nil {
		nextHandler, nextMeta, err = p.rx.FirstHandler(handler.SortASC)
	} else if p.rx.HasNext(*message.meta, handler.SortASC) {
		nextHandler, nextMeta, err = p.rx.NextHandler(*message.meta)
	} else {
		fmt.Printf("没有handler了,打印Payload:%v", message.payload)
		return
	}

	if err != nil {
		return
	}
	err = nextHandler.Rx(message.payload)
	if err != nil {
		return
	}
	var n = 0
	var context handler.Context = nextHandler
	if context != nil {
		n = context.RxLen()
	}
	for i := 0; i < n; i++ {
		var r interface{}
		r, err = nextHandler.RxRead()
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

func (p *Pipeline) OnTx(data interface{}) {
	var message = MessageContext{
		payload: data,
		meta:    nil,
	}
	err := p.handleTx(message)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Pipeline) handleTx(message MessageContext) (err error) {

	var nextHandler handler.Handler
	var nextMeta handler.Meta
	if message.meta == nil {
		nextHandler, nextMeta, err = p.tx.FirstHandler(handler.SortDESC)
	} else if p.tx.HasNext(*message.meta, handler.SortDESC) {
		nextHandler, nextMeta, err = p.tx.NextHandler(*message.meta)
	} else {
		return
	}

	if err != nil {
		return
	}
	err = nextHandler.Tx(message.payload)

	if err != nil {
		return
	}
	var n = 0
	var context handler.Context = nextHandler
	if context != nil {
		n = context.TxLen()
	}
	for i := 0; i < n; i++ {
		var r interface{}
		r, err = nextHandler.TxRead()
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
