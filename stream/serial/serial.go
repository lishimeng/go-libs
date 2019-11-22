package serial

import (
	"github.com/lishimeng/go-libs/log"
	delegate "github.com/tarm/goserial"
	"io"
)

type Connector struct {
	// 码率
	BaudRate uint

	// 串口
	PortName string

	Ser io.ReadWriteCloser
}

func New(rate uint, name string) *Connector {

	return &Connector{
		BaudRate: rate,
		PortName: name,
	}
}

func (c *Connector) Connect() (err error) {
	log.Fine("open connect: %s[%d]", c.PortName, c.BaudRate)
	serialOptions := &delegate.Config{
		Name: c.PortName,
		Baud: int(c.BaudRate),
	}
	c.Ser, err = delegate.OpenPort(serialOptions)
	if err != nil {
		log.Fine("can't connect to: %s[%d], %s", c.PortName, c.BaudRate, err.Error())
		return
	}
	log.Fine("connect success")
	return
}

func (c Connector) Close() error {
	log.Fine("close the connection: %s", c.PortName)
	return c.Ser.Close()
}
