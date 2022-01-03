package app

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/jszwec/csvutil"
)

func HeaderFromString(v interface{}) (header []string, err error) {
	return csvutil.Header(v, "csv")
}

type App struct {
	r *csv.Reader
	w *csv.Writer
	d *csvutil.Decoder
	e *csvutil.Encoder

	err error
}

// TODO: need to prettify
type Options struct {
	SetHeader bool
	UseCRLF   bool
	Comma     rune
	Comment   rune
	Schema    interface{}
}

func New(w io.Writer, r io.Reader, options ...*Options) (*App, error) {
	var h []string
	a := &App{}

	if options == nil {
		options = append(options, &Options{})
	}

	if options[0].SetHeader {
		h, a.err = HeaderFromString(options[0].Schema)
		if a.err != nil {
			return nil, a.err
		}
	}

	a.e = csvutil.NewEncoder(a.newWriter(w, options[0]))
	a.d, a.err = csvutil.NewDecoder(a.newReader(r, options[0]), h...)
	return a, a.err
}

func (a *App) Flush() error               { a.w.Flush(); return a.w.Error() }
func (a *App) Error() error               { return a.w.Error() }
func (a *App) Encode(v interface{}) error { return a.e.Encode(v) }
func (a *App) Header() []string           { return a.d.Header() }
func (a *App) Scan() bool                 { return a.err == nil }
func (a *App) Decode(v interface{}) error {
	a.err = a.d.Decode(&v)
	return a.err
}

func (a *App) Print() {
	fmt.Printf("a.w: %v\n", a.w)
}

func (a *App) newWriter(w io.Writer, options *Options) *csv.Writer {
	a.w = csv.NewWriter(w)
	if options != nil {
		if options.Comma != 0 {
			a.w.Comma = options.Comma
		}
	}
	return a.w
}

func (a *App) newReader(r io.Reader, options *Options) *csv.Reader {
	a.r = csv.NewReader(r)
	if options != nil {
		if options.Comma != 0 {
			a.r.Comma = options.Comma
		}
		if options.Comment != 0 {
			a.r.Comment = options.Comment
		}
	}
	return a.r
}
