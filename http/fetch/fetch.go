package fetch

import (
	"net/http"
)

func Fetch(url string, callback func(resp *http.Response) error, o *Options) error {
	req, err := http.NewRequest(o.m, url, o.b)
	if err != nil {
		return err
	}
	req.Header = o.h

	resp, err := o.cli.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return callback(resp)
}
