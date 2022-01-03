package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/topheruk/go/learn/data/csv/model"
	"github.com/topheruk/go/learn/data/csv/serde"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type app struct {
	csv map[string]*serde.CSV
}

func run() (err error) {
	fr, err := os.Open("learn/data/csv/data/accounts.csv")
	if err != nil {
		return
	}
	defer fr.Close()

	var buf bytes.Buffer
	// TODO: time should be in format ISO8601:YYYY-MM-DDTHH:MM:SSZ
	// FIXME: all my csvs have a value for canvas_*_id which does not need to be there
	csv, err := serde.NewCSV(&buf, fr, &serde.CSVOptions{TimeFormat: "0001-01-01T00:00:00Z"})
	if err != nil {
		return
	}

	sd := model.AccountSerde{Serde: csv}

	v, err := sd.Get()
	if err != nil {
		return
	}
	sd.Set(v)

	fmt.Println(buf.String())

	return nil
}
