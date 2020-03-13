package serial

import (
	delegate "github.com/tarm/serial"
	"io"
)

type Config delegate.Config

type Connector struct {
	config *Config

	Ser io.ReadWriteCloser
}

func New(config *Config) *Connector {

	return &Connector{
		config: config,
	}
}

func (c *Connector) Connect() (err error) {
	var conf = (*delegate.Config)(c.config)
	c.Ser, err = delegate.OpenPort(conf)
	if err != nil {
		return
	}
	return
}

func (c Connector) Close() error {
	return c.Ser.Close()
}
