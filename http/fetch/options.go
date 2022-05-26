package fetch

import (
	"io"
	"net/http"
	"strings"
)

type options struct {
	cli *http.Client
	m   string //method
	h   http.Header
	b   io.Reader
}

func optionsBuilder() *options {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")

	return &options{
		cli: http.DefaultClient,
		h:   h,
		m:   "GET",
	}
}

var DefaultOptions = optionsBuilder()

func (o *options) Client(cli *http.Client) *options {
	o.cli = cli
	return o
}

func (o *options) ContentType(contentTyp string) *options {
	o.h.Set("Content-Type", contentTyp)
	return o
}

func (o *options) Header(h http.Header) *options {
	for k, v := range h {
		o.h.Set(k, strings.Join(v, "; "))
	}
	return o
}
