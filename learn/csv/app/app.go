package app

import (
	"encoding/csv"
	"os"
)

type App struct {
	r      *csv.Reader
	err    error
	record []string
}

func New(f *os.File) (a *App) {
	a = &App{r: csv.NewReader(f)}
	return
}

func (a *App) Scan() bool {
	a.record, a.err = a.r.Read()
	return a.err == nil
}

func (a *App) Record() (record []string)           { return a.record }
func (a *App) All() (record [][]string, err error) { return a.r.ReadAll() }

func (a *App) Encode() (err error) {

	return
}
