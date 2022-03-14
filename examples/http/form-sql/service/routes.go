package service

//
//

func (s *Service) routes() {
	s.m.Get("/pong", s.Echo("pong"))
}
