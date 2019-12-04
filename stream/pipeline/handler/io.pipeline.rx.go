package handler

func (c *RxContainer) NextHandler(currentMeta Meta) (handler Handler, meta Meta, ctx *Context, err error) {
	handler, meta, ctx, err = c.Next(currentMeta, SortASC)

	if err == nil {
		if meta.DisableRx {
			return c.NextHandler(meta)
		}
	}
	return
}
