package sync

type Trigger interface {
	Trigger()
}

func (s *Service) Trigger() {
	s.trigger <- struct{}{}
}
