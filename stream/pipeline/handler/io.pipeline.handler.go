package handler

import "fmt"

type Handler interface {
	Context
	Rx(input interface{}) (err error)
	Tx(input interface{}) (err error)
}

type Context interface {
	RxLen() int
	TxLen() int
	RxWrite(data interface{})
	RxRead() (data interface{}, err error)
	TxWrite(data interface{})
	TxRead() (data interface{}, err error)
}

type ChanBasedContext struct {
	rx     chan interface{}
	tx     chan interface{}
	txSize int
	rxSize int
}

func NewContext(rxCapacity int, txCapacity int) *Context {

	c := &ChanBasedContext{
		rx:     make(chan interface{}, rxCapacity),
		tx:     make(chan interface{}, txCapacity),
		rxSize: 0,
		txSize: 0,
	}
	var inter Context = c
	return &inter
}

func (ccb *ChanBasedContext) RxLen() int {
	return ccb.rxSize
}

func (ccb *ChanBasedContext) TxLen() int {
	return ccb.txSize
}

func (ccb *ChanBasedContext) RxWrite(data interface{}) {
	ccb.rx <- data
	ccb.rxSize++
}

func (ccb *ChanBasedContext) RxRead() (data interface{}, err error) {
	if ccb.rxSize > 0 {
		data = <-ccb.rx
		ccb.rxSize--
	} else {
		err = fmt.Errorf("o length")
	}
	return
}

func (ccb *ChanBasedContext) TxWrite(data interface{}) {
	ccb.tx <- data
	ccb.txSize++
}

func (ccb *ChanBasedContext) TxRead() (data interface{}, err error) {
	if ccb.txSize > 0 {
		data = <-ccb.tx
		ccb.txSize--
	} else {
		err = fmt.Errorf("o length")
	}
	return
}
