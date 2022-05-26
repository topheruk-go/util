package fetch

import (
	"io"
	"net/http"
	"strings"
)

type Options struct {
	cli *http.Client
	m   string //method
	h   http.Header
	b   io.Reader
}

var DefaultOptions = &Options{
	cli: http.DefaultClient,
	h: http.Header{
		"Content-Type": []string{"appliaction/json"},
	},
	m: "GET",
}

func (o *Options) Client(cli *http.Client) *Options {
	o.cli = cli
	return o
}

func (o *Options) ContentType(contentTyp string) *Options {
	o.h.Set("Content-Type", contentTyp)
	return o
}

func (o *Options) Header(h http.Header) *Options {
	for k, v := range h {
		o.h.Set(k, strings.Join(v, "; "))
	}
	return o
}
