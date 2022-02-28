package service

func (s *Service) routes() {
	s.m.Get("/ping", s.Echo("ping"))
}
