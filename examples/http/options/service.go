package options

import "time"

type service struct {
	address        string
	timeout        time.Duration
	maxConnections int
}

type option func(*service)

func newService(address string, os ...option) *service {
	service := &service{
		address: address,
	}

	for _, o := range os {
		o(service)
	}

	return service
}

func withTimeOut(t time.Duration) option {
	return func(s *service) {
		s.timeout = t
	}
}

func withMaxConn(m int) option {
	return func(s *service) {
		s.maxConnections = m
	}
}
