package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/topheruk/go/learn/net/mongodb/app"
	"github.com/topheruk/go/learn/net/mongodb/database"
)

type TestCase struct {
	desc        string
	method      string
	path        string
	contentType string
	content     string
	code        int
}

var tcs = []TestCase{
	{desc: "add to bar collection", method: "POST", path: "/api/v1/bar/", contentType: "application/json", content: `{ "value":"Value" }`, code: http.StatusOK},
	{desc: "add to foo collection", method: "POST", path: "/api/v1/foo/", contentType: "application/json", content: `{ "value":100 }`, code: http.StatusOK},
	// delete methods required when I figure how to have custom IDs
}

var user = flag.String("user", "", "client username")
var pass = flag.String("pass", "", "client username")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestRequest(t *testing.T) {
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", *user, *pass)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "foobar")
	if err != nil {
		return
	}
	defer db.Client().Disconnect(ctx)

	srv := httptest.NewServer(app.New(db))
	defer srv.Close()

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			if err = run(srv, &tc); err != nil {
				t.Error(err)
			}
		})
	}
}

func run(srv *httptest.Server, tc *TestCase) (err error) {
	q, err := http.NewRequest(tc.method, srv.URL+tc.path, bytes.NewBufferString(tc.content))
	q.Header.Set("Content-Type", tc.contentType)
	if err != nil {
		return
	}

	r, err := srv.Client().Do(q)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != tc.code {
		return fmt.Errorf("expected %v; got %v", tc.code, r.Status)
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
