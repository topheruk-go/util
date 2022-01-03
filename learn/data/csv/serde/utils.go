package serde

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
)

func HeaderFromString(v interface{}) (header []string, err error) {
	return csvutil.Header(v, "csv")
}

func newEncoder(w *csv.Writer) *csvutil.Encoder {
	return csvutil.NewEncoder(w)
}

func newDecoder(r *csv.Reader, options *CSVOptions) (*csvutil.Decoder, error) {
	var h []string
	var err error
	if options != nil && options.SetHeader {
		h, err = HeaderFromString(options.Schema)
		if err != nil {
			return nil, err
		}
	}
	return csvutil.NewDecoder(r, h...)
}

func newWriter(w io.Writer, options *CSVOptions) *csv.Writer {
	cw := csv.NewWriter(w)
	if options != nil {
		if options.Comma != 0 {
			cw.Comma = options.Comma
		}
		if options.UseCRLF {
			cw.UseCRLF = options.UseCRLF
		}
	}
	return cw
}

func newReader(r io.Reader, options *CSVOptions) *csv.Reader {
	cr := csv.NewReader(r)
	if options != nil {
		if options.Comma != 0 {
			cr.Comma = options.Comma
		}
		if options.Comment != 0 {
			cr.Comment = options.Comment
		}
	}
	return cr
}
