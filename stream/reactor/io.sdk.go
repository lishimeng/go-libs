package reactor

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/lishimeng/go-libs/log"
	"io"
	"time"
)

const (
	ModeSync = iota
	ModeAsync
)

var (
	TimeoutErr = errors.New("timeout")
)

type Stream struct {
	io.Reader
	io.Writer
	io.Closer

	// sync
	buf      *bytes.Buffer
	syncChan chan bool

	// async
	dataListener func(bytes []byte)

	rxChan    chan []byte
	txChan    chan []byte
	closeChan chan byte

	onLostConnect func(err error)

	mode int
}

func New(r io.ReadWriteCloser) *Stream {
	return CreateWithRWC(r, r, r)
}

func CreateWithRWC(r io.Reader, w io.Writer, c io.Closer) *Stream {

	s := &Stream{
		rxChan:    make(chan []byte, 1024),
		txChan:    make(chan []byte, 1024),
		closeChan: make(chan byte),
		buf:       bytes.NewBuffer(make([]byte, 4096)),
		syncChan:  make(chan bool),
		mode:      ModeAsync,
	}
	s.buf.Reset()
	s.Reader = r
	s.Writer = w
	s.Closer = c

	go s.readLoop()
	go s.writeLoop()
	go s.processLoop()

	return s
}

func (s *Stream) DataListener(listener func([]byte)) {
	s.dataListener = listener
	s.mode = ModeAsync
}

func (s *Stream) Write(p []byte) (n int, err error) {
	return s.Writer.Write(p)
}

/*
阻塞读取
p 数据
timeout 超时间隔
err 超时返回
*/
func (s *Stream) ReadSync(p []byte, timeout time.Duration) (err error) {

	s.mode = ModeSync
	var timer *time.Timer
	// 直接读
	if s.buf.Len() >= len(p) {
		log.Fine("直接读:%d", len(p))
		_, err = s.buf.Read(p)
		return
	}
	timer = time.NewTimer(timeout)
	log.Fine("阻塞读:%d[%d]", len(p), timeout)
	defer func() {
		if timer != nil {
			log.Fine("撤销timeout")
			timer.Stop()
		}
	}()
	select {
	case <-s.closeChan:
		break
	case <-timer.C:
		err = TimeoutErr
		// TODO
	}
	return
}

func (s *Stream) Close() {
	log.Fine("close chan")
	close(s.closeChan)
	log.Fine("wait 2000ms")
	time.Sleep(time.Millisecond * 2000)
	if s.Closer != nil {
		err := s.Closer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
