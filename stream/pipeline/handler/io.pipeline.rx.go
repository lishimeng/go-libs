package handler

func (c *RxContainer) NextHandler(currentMeta Meta) (handler Handler, meta Meta, err error) {
	handler, meta, err = c.Next(currentMeta, SortASC)

	if err == nil {
		if meta.DisableRx {
			return c.NextHandler(meta)
		}
	}
	return
}
