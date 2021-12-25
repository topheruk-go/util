package template

import (
	"bytes"
	"html/template"
	"net/http"
	"sync"
)

func Render(path ...string) (f func(rw http.ResponseWriter, r *http.Request, data interface{}), err error) {
	var (
		init sync.Once

		tpl *template.Template
		buf *bytes.Buffer
	)

	init.Do(func() { tpl, err = template.ParseFiles(path...) })

	return func(rw http.ResponseWriter, r *http.Request, data interface{}) {
		buf = &bytes.Buffer{}
		err = tpl.Execute(buf, data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(rw)
	}, err
}
