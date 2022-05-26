package fetch

import (
	"io"
	"net/http"
)

type options struct {
	cli *http.Client
	m   string      //method
	h   http.Header //applicationTyp
	b   io.Reader
}

func optionsBuilder() *options {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")

	return &options{
		cli: http.DefaultClient,
		h:   h,
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

func Fetch(url string, callback func(resp *http.Response) error, o *options) error {
	req, err := http.NewRequest(o.m, url, o.b)
	if err != nil {
		return err
	}

	resp, err := o.cli.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return callback(resp)
}
