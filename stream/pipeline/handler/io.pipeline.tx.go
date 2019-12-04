package handler

func (c *TxContainer) NextHandler(currentMeta Meta) (handler Handler, meta Meta, ctx *Context, err error) {
	handler, meta, ctx, err = c.Next(currentMeta, SortDESC)

	if err == nil {
		if meta.DisableTx {
			return c.NextHandler(meta)
		}
	}
	return
}
