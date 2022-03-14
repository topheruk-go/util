package options

import "time"

type service struct {
	address        string
	timeout        time.Duration
	maxConnections int
}

type option func(*service)

type builder struct {
	opts []option
}

func newBuilder() *builder {
	b := &builder{
		make([]option, 0),
	}
	return b
}

func (b *builder) withTimeOut(t time.Duration) *builder {
	b.opts = append(b.opts, withTimeOut(t))
	return b
}

func (b *builder) withMaxConn(m int) *builder {
	b.opts = append(b.opts, withMaxConn(m))
	return b
}

func (b *builder) build() []option {
	return b.opts
}

// -- options builder
// this struct would allow me to do the following:
// opts := Options().WithTimeout(time_dureation).WithMaxConn(max_conn).Build()

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
