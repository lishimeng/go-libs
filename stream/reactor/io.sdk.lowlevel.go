package reactor

import (
	"io"
)

func (s *Stream) readLoop() {
	defer func() {

	}()

	for {
		select {
		case <-s.closeChan:
			return
		case p := <-s.rxChan:
			if s.dataListener != nil {
				s.dataListener(p)
			}
		}
	}
}

func (s *Stream) writeLoop() {

	defer func() {

	}()

	for {
		select {
		case <-s.closeChan:
			return
		case p := <-s.txChan:
			_, err := s.Writer.Write(p)
			if err != nil {
				s.lostConnection(err)
			}
		}
	}
}

func (s *Stream) processLoop() {

	defer func() {
		close(s.rxChan)
		close(s.txChan)
	}()
	buf := make([]byte, 1024)
	for {
		n, err := s.Reader.Read(buf)

		if err != nil {
			if err != io.EOF { // has error
				s.lostConnection(err)
			}
			break
		}
		if n > 0 {
			data := buf[:n]
			s.buf.Write(data)

			switch s.mode {
			case ModeAsync:
				// copy
				p := s.buf.Next(s.buf.Len())
				s.rxChan <- p
			case ModeSync:
				s.syncChan <- true
			}
		} else {
		}
	}
}
