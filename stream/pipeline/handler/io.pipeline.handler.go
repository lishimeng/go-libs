package handler

import "fmt"

type Handler interface {
	Rx(input interface{}, ctx *Context) (err error)
	Tx(input interface{}, ctx *Context) (err error)
}

type Context struct {
	rx     chan interface{}
	tx     chan interface{}
	txSize int
	rxSize int
}

func NewContext(rxCapacity int, txCapacity int) *Context {

	c := &Context{
		rx:     make(chan interface{}, rxCapacity),
		tx:     make(chan interface{}, txCapacity),
		rxSize: 0,
		txSize: 0,
	}
	return c
}

func (ccb *Context) RxLen() int {
	return ccb.rxSize
}

func (ccb *Context) TxLen() int {
	return ccb.txSize
}

func (ccb *Context) RxWrite(data interface{}) {
	ccb.rx <- data
	ccb.rxSize++
}

func (ccb *Context) RxRead() (data interface{}, err error) {
	if ccb.rxSize > 0 {
		data = <-ccb.rx
		ccb.rxSize--
	} else {
		err = fmt.Errorf("o length")
	}
	return
}

func (ccb *Context) TxWrite(data interface{}) {
	ccb.tx <- data
	ccb.txSize++
}

func (ccb *Context) TxRead() (data interface{}, err error) {
	if ccb.txSize > 0 {
		data = <-ccb.tx
		ccb.txSize--
	} else {
		err = fmt.Errorf("o length")
	}
	return
}
