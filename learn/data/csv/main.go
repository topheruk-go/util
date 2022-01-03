package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

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
	csv map[string]*serde.Serde
}

func run() (err error) {
	fr, err := readFile("learn/data/csv/data/accounts.csv")
	if err != nil {
		return
	}
	defer fr.Close()

	var buf bytes.Buffer
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

	// read directory
	err = readDir("./learn/data/csv/data")
	if err != nil {
		return
	}

	return nil
}

func readFile(name string) (*os.File, error) {
	return os.Open(name)
}

func readDir(name string) (err error) {
	d, err := os.Open(name)
	if err != nil {
		return
	}

	defer d.Close()

	list, err := d.Readdir(-1)
	if err != nil {
		return
	}

	for _, f := range list {
		switch {
		case filepath.Ext(f.Name()) == ".csv":
			if err := newSerde(name + "/" + f.Name()); err != nil {
				return err
			}
		}

	}

	return
}

func newSerde(name string) error {
	fmt.Printf("%s\n", name)
	return nil
}
