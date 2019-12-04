package handler

func (c *TxContainer) NextHandler(currentMeta Meta) (handler Handler, meta Meta, err error) {
	handler, meta, err = c.Next(currentMeta, SortASC)

	if err == nil {
		if meta.DisableTx {
			return c.NextHandler(meta)
		}
	}
	return
}
