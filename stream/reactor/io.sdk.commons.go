package reactor

func (s *Stream) lostConnection(err error) {
	if s.onLostConnect != nil {
		s.onLostConnect(err)
	}
}
