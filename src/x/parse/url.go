package parse

import (
	"net/http"
	"strings"
)

func AbsoluteURL(rw http.ResponseWriter, r *http.Request) string {
	return strings.ToLower(strings.SplitN(r.Proto, "/", 2)[0]) + "://" + r.Host + r.URL.String()
	// return fmt.Sprintf("%s://%s%s", strings.ToLower(strings.SplitN(r.Proto, "/", 2)[0]), r.Host, r.URL)
}
