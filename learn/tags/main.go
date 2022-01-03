package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/topheruk/go/learn/data/csv/app"
)

var filename = flag.String("f", "", "name given for .csv file")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() (err error) {
	f, err := os.Open(fmt.Sprintf("learn/data/csv/data/%s.csv", *filename))
	if err != nil {
		return
	}
	defer f.Close()

	app := app.New(f)

	for app.Scan() {
		for key := range app.Record() {
			fmt.Printf("%s\n", app.Record()[key])
		}
	}

	var u User
	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Tag: %s, TypeName: %s Value: %v\n",
			field.Name, field.Tag, field.Type.Name(), value)
	}

	fmt.Println(t)

	var user User
	var records = []string{"John", "32", "London"}
	Decode(records, &user)

	fmt.Println(user)

	return
}

// panics if interface is not a pointer to it
// but should just return an error & not do anything to the interface that was passed
func Decode(record []string, v interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(v).Elem())
	for i, r := range record {
		f := rv.Field(i)
		_switch(f, r)
	}
}

func _switch(f reflect.Value, s string) {
	if f.IsValid() && f.CanSet() {
		switch k := f.Kind(); {
		case k == reflect.String:
			f.SetString(s)
		case k == reflect.Int:
			x, _ := strconv.Atoi(s)
			f.SetInt(int64(x))
		}
	}
}

type User struct {
	Name     string `csv:"name"`
	Age      int    `csv:"age"`
	Location string `csv:"location"`
}

package app

import (
	"encoding/csv"
	"os"
)

type App struct {
	r      *csv.Reader
	
}

func New(f *os.File) (a *App) {
	a = &App{r: csv.NewReader(f)}
	return
}

func (a *App) Scan() bool {
	a.record, a.err = a.r.Read()
	return a.err == nil
}

func (a *App) Record() []string { return a.record }
