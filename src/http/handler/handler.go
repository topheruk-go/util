package handler

import (
	"net/http"

	"github.com/topheruk/go/src/encoding"
)

func Echo(message string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) { encoding.Respond(rw, r, message, http.StatusOK) }
}

func FileServer(prefix, dirname string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dirname)))
}
