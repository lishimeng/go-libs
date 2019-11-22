package reactor

import (
	"github.com/lishimeng/go-libs/log"
	"io"
)

func (s *Stream) readLoop() {
	log.Fine("start read loop")
	defer func() {

	}()

	for {
		select {
		case <-s.closeChan:
			log.Fine("stop read loop")
			return
		case p := <-s.rxChan:
			log.Fine("read loop received:%s(%d)", string(p), len(p))
			if s.dataListener != nil {
				s.dataListener(p)
			}

			log.Fine("--------------------------------------------------")
		}
	}
}

func (s *Stream) writeLoop() {

	log.Fine("start write loop")
	defer func() {

	}()

	for {
		select {
		case <-s.closeChan:
			log.Fine("stop write loop")
			return
		case p := <-s.txChan:
			log.Fine("write loop send:-->")
			n, err := s.Writer.Write(p)
			if err != nil {
				s.lostConnection(err)
			}
			log.Fine("write size:%d", n)
		}
	}
}

func (s *Stream) processLoop() {

	log.Fine("start process loop")
	defer func() {
		log.Fine("close rx chan")
		close(s.rxChan)
		log.Fine("close tx chan")
		close(s.txChan)
	}()
	buf := make([]byte, 1024)
	for {
		log.Fine("read:<---")
		n, err := s.Reader.Read(buf)

		if err != nil {
			if err != io.EOF { // has error
				log.Fine("read error:%v", err)
				s.lostConnection(err)
			}
			break
		}
		if n > 0 {
			data := buf[:n]
			s.buf.Write(data)
			log.Fine("write to buf:(%d)%s", n, string(data))

			switch s.mode {
			case ModeAsync:
				// copy
				p := s.buf.Next(s.buf.Len())
				s.rxChan <- p
			case ModeSync:
				s.syncChan <- true
			}
		} else {
			log.Fine("读到%d个", n)
		}
	}
}
