package serial

import (
	"github.com/lishimeng/go-libs/log"
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
	log.Fine("open connect: %s[%d]", c.config.Name, c.config.Baud)
	var conf = (*delegate.Config)(c.config)
	c.Ser, err = delegate.OpenPort(conf)
	if err != nil {
		log.Fine("can't connect to: %s[%d], %s", c.config.Name, c.config.Baud, err.Error())
		return
	}
	log.Fine("connect success")
	return
}

func (c Connector) Close() error {
	log.Fine("close the connection: %s", c.config.Name)
	return c.Ser.Close()
}
