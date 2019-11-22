package serialnet

import (
	serial2 "github.com/lishimeng/go-libs/serial/serial"
)

type Config struct {
	PortName string
	BondRate uint
	NetPort  uint
}

type Connector struct {
	serial2.Connector
}

func New(conf Config) *Connector {

	return nil
}
