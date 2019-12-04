package handler

import (
	"fmt"
	"time"
)

func NewContainer() (container *Container) {
	container = &Container{
		coders: make(map[string]Handler),
		metas:  make(map[string]Meta),
	}
	return container
}

func (c *Container) RegisterHandlers(handlers ...Handler) (err error) {

	if len(handlers) > 0 {
		for _, h := range handlers {
			err = c.Register(h)
			if err != nil {
				break
			}
		}
	}
	return
}

func (c *Container) Register(handler Handler, meta ...Meta) (err error) {

	if handler == nil {
		err = fmt.Errorf("object nil")
		return
	}

	if !checkHandler(handler) {
		err = fmt.Errorf("object must implement Context")
		return
	}

	currentHandlerSize := len(c.ens)
	name := fmt.Sprintf("coder_%d_%d", currentHandlerSize, time.Now().UnixNano())
	index := len(c.ens)
	c.coders[name] = handler
	c.ens = append(c.ens, name)
	var m Meta
	if len(meta) > 0 {
		m = meta[0]
	} else {
		m = Meta{
			DisableRx: false,
			DisableTx: false,
		}
	}
	m.index = index
	c.metas[name] = m
	return
}

func checkHandler(handler Handler) (b bool) {

	defer func() {
		if err := recover(); err != nil {
			b = false
		}
	}()
	handler.TxLen()
	b = true
	return
}

func (c Container) HasNext(meta Meta, sort CoderSort) bool {

	nextIndex := calcNextIndex(meta.index, sort)
	return validIndex(nextIndex, len(c.ens))
}

func validIndex(index int, cap int) bool {

	return index >= 0 && index < cap
}

func calcNextIndex(current int, sort CoderSort) (next int) {

	if sort == SortASC {
		next = current + 1
	} else {
		next = current - 1
	}
	return
}

func (c *Container) Get(name string) (handler Handler, meta Meta, err error) {

	handler, ok := c.coders[name]
	if !ok {
		err = fmt.Errorf("no codec named:%s", name)
	}
	if ok {
		meta, ok = c.metas[name]
		if !ok {
			err = fmt.Errorf("no codec meta named:%s", name)
		}
	}
	return
}

func (c *Container) Next(currentMeta Meta, sort CoderSort) (handler Handler, meta Meta, err error) {
	// calc next index
	var nextIndex = calcNextIndex(currentMeta.index, sort)
	if !validIndex(nextIndex, len(c.ens)) {
		err = fmt.Errorf("there is no more codec exist")
		return
	}
	return c.getHandler(nextIndex)
}

func (c *Container) FirstHandler(sort CoderSort) (handler Handler, meta Meta, err error) {
	// 0--len(coders)
	// calc next index
	if len(c.ens) == 0 {
		err = fmt.Errorf("there is no more codec exist")
		return
	}
	var nextIndex int
	if sort == SortASC {
		nextIndex = 0
	} else {
		nextIndex = len(c.ens) - 1 // 最后一个
	}
	return c.getHandler(nextIndex)
}

func (c *Container) getHandler(index int) (handler Handler, meta Meta, err error) {
	name := c.ens[index]
	return c.Get(name)
}
