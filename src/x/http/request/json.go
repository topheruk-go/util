package request

import (
	"net/http"

	"github.com/topheruk/go/src/x/encoding"
)

func New(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	return encoding.Decode(rw, r, data)
}
