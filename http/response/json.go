package response

import (
	"errors"
	"net/http"

	"github.com/topheruk/go/encoding"
)

func New(rw http.ResponseWriter, r *http.Request, data interface{}, status int) {
	encoding.Respond(rw, r, data, status)
}

func Err(rw http.ResponseWriter, r *http.Request, err error, status int) {
	type e struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	New(rw, r, e{status, err.Error()}, status)
}

func Todo(rw http.ResponseWriter, r *http.Request) {
	Err(rw, r, errors.New("todo"), http.StatusInternalServerError)
}
