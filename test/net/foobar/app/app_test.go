package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/topheruk/go/test/net/foobar/app"
	"github.com/topheruk/go/test/net/foobar/database"
)

type request struct {
	desc        string
	method      string
	path        string
	contentType string
	content     string
	code        int
}

var qs = []request{
	{desc: "add to bar collection", method: "POST", path: "/api/v1/bar/", contentType: "application/json", content: `{ "value":"Value" }`, code: http.StatusOK},
	{desc: "add to foo collection", method: "POST", path: "/api/v1/foo/", contentType: "application/json", content: `{ "value":100 }`, code: http.StatusOK},
}

// var user, pass *string

// func TestMain(m *testing.M) {
// 	user, pass, _ = parse.Flags()
// 	os.Exit(m.Run())
// }

func TestRequest(t *testing.T) {
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", "topheruk", "T^*G7!Pf")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "foobar")
	if err != nil {
		return
	}
	defer db.Client().Disconnect(ctx)

	srv := httptest.NewServer(app.New(db))
	defer srv.Close()

	for _, q := range qs {
		t.Run(q.desc, func(t *testing.T) {
			if err = NewRequest(srv, &q); err != nil {
				t.Error(err)
			}
		})
	}
}

func NewRequest(srv *httptest.Server, t *request) (err error) {
	q, err := http.NewRequest(t.method, srv.URL+t.path, bytes.NewBufferString(t.content))
	q.Header.Set("Content-Type", t.contentType)
	if err != nil {
		return
	}

	r, err := srv.Client().Do(q)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != t.code {
		return fmt.Errorf("expected %v; got %v", t.code, r.Status)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	if r.StatusCode == http.StatusOK {
		var v interface{}
		if err := json.Unmarshal(b, &v); err != nil {
			return fmt.Errorf("could not unmarshal body: %v", err)
		}
	}

	return
}
