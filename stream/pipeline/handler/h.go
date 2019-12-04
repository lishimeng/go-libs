package handler

type CoderSort int

const (
	SortASC CoderSort = iota
	SortDESC
)

// define
type Meta struct {
	DisableTx bool
	DisableRx bool
	index     int
}

func (c Meta) Index() int {
	return c.index
}

type Container struct {
	coders map[string]Handler
	metas  map[string]Meta
	ens    []string
}

type TxContainer struct {
	Container
}
type RxContainer struct {
	Container
}
