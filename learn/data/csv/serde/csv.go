package serde

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

func HeaderFromString(v interface{}) (header []string, err error) {
	return csvutil.Header(v, "csv")
}

type CSV struct {
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

func NewCSV(w io.Writer, r io.Reader, options ...*Options) (*CSV, error) {
	var h = []string{}
	var a = &CSV{}

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

func (a *CSV) Flush() error               { a.w.Flush(); return a.w.Error() }
func (a *CSV) Err() error                 { return a.err }
func (a *CSV) Encode(v interface{}) error { return a.e.Encode(v) }
func (a *CSV) Header() []string           { return a.d.Header() }
func (a *CSV) Scan() bool                 { return a.err == nil }
func (a *CSV) Decode(v interface{}) error { a.err = a.d.Decode(&v); return a.Err() }

func (a *CSV) newWriter(w io.Writer, options *Options) *csv.Writer {
	a.w = csv.NewWriter(w)
	if options != nil {
		if options.Comma != 0 {
			a.w.Comma = options.Comma
		}
		if options.UseCRLF {
			a.w.UseCRLF = options.UseCRLF
		}
	}
	return a.w
}

func (a *CSV) newReader(r io.Reader, options *Options) *csv.Reader {
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

type Store interface {
	Decode(v interface{}) error
	Encode(v interface{}) error
	Scan() bool
}
