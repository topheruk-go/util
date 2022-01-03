package app

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

func HeaderFromString(v interface{}) (header []string, err error) {
	return csvutil.Header(v, "csv")
}

type App struct {
	d *csvutil.Decoder
	e *csvutil.Encoder

	err error
}

// TODO: need to prettify
type AppOptions struct {
	SetHeader bool
	Comma     rune
	Comment   rune
	Schema    interface{}
}

func New(w io.Writer, r io.Reader, options ...*AppOptions) (a *App, err error) {
	var h []string
	if options == nil {
		options = append(options, &AppOptions{})
	}

	if options[0].SetHeader {
		h, err = HeaderFromString(options[0].Schema)
		if err != nil {
			return nil, err
		}
	}

	a = &App{
		e: csvutil.NewEncoder(newWriter(w, options[0])),
	}

	a.d, err = csvutil.NewDecoder(newReader(r, options[0]), h...)
	return
}

func (a *App) Err() error                 { return a.err }
func (a *App) Encode(v interface{}) error { return a.e.Encode(v) }
func (a *App) Header() []string           { return a.d.Header() }
func (a *App) Scan() bool                 { return a.err == nil }
func (a *App) Decode(v interface{}) error {
	a.err = a.d.Decode(&v)
	return a.err
}

func newWriter(w io.Writer, options *AppOptions) *csv.Writer {
	csvw := csv.NewWriter(w)
	if options != nil {
		if options.Comma != 0 {
			csvw.Comma = options.Comma
		}
	}
	return csvw
}

func newReader(r io.Reader, options *AppOptions) *csv.Reader {
	csvr := csv.NewReader(r)
	if options != nil {
		if options.Comma != 0 {
			csvr.Comma = options.Comma
		}
		if options.Comment != 0 {
			csvr.Comment = options.Comment
		}
	}
	return csvr
}
