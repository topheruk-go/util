package serde

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

type CSV struct {
	r *csv.Reader
	w *csv.Writer
	d *csvutil.Decoder
	e *csvutil.Encoder

	err error
}

// TODO: need to prettify
type CSVOptions struct {
	SetHeader bool
	UseCRLF   bool
	Comma     rune
	Comment   rune
	Schema    interface{}
	// TimeFormat string
}

func NewCSV(w io.Writer, r io.Reader, options *CSVOptions) (c *CSV, err error) {
	c = &CSV{
		w: newWriter(w, options),
		r: newReader(r, options),
	}

	if c.d, err = newDecoder(c.r, options); err != nil {
		return
	}

	c.e = newEncoder(c.w)
	return c, nil
}

func (a *CSV) Flush() error                                               { a.w.Flush(); return a.w.Error() }
func (a *CSV) Err() error                                                 { return a.err }
func (a *CSV) Encode(v interface{}) error                                 { return a.e.Encode(v) }
func (a *CSV) Header() []string                                           { return a.d.Header() }
func (a *CSV) Decode(v interface{}) error                                 { a.err = a.d.Decode(&v); return a.Err() }
func (a *CSV) Map(f func(field string, col string, v interface{}) string) { a.d.Map = f }
