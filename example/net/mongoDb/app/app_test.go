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

	"github.com/topheruk/go/example/net/mongoDb/app"
	"github.com/topheruk/go/example/net/mongoDb/database"
)

var (
	username = ""
	password = ""
	host     = ""
	port     = 27017
	uri      = fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
)

type requestTest struct {
	desc        string
	method      string
	path        string
	code        int
	contentType string
	content     string
}

var requests = []requestTest{
	{desc: "all users in databse found", method: "GET", path: "/api/v1/users/", contentType: "application/json", code: http.StatusOK},
	{desc: "redirecting if trailing slash exists", method: "GET", path: "/api/v1/users", contentType: "application/json", code: http.StatusNotFound},
	{desc: "user found in database", method: "GET", path: "/api/v1/users/61ce33e3928e6155964a629f", contentType: "application/json", code: http.StatusOK},
	{desc: "requesting with invalid id", method: "GET", path: "/api/v1/users/34ce33e5964a629f", contentType: "application/json", code: http.StatusBadRequest},
	{desc: "no user with valid id", method: "GET", path: "/api/v1/users/34ce33e3928e6155964a629f", contentType: "application/json", code: http.StatusInternalServerError},
	{desc: "creating a new user", method: "POST", path: "/api/v1/users/", contentType: "application/json", content: `{ "name":"Justin", "age":15 }`, code: http.StatusOK},
	{desc: "invalid create user request", method: "POST", path: "/api/v1/users/", contentType: "application/json", content: `{ "name":"Justin" }`, code: http.StatusBadRequest},
}

var createRequests = []requestTest{
	{desc: "creating a new user", method: "POST", path: "/api/v1/users/", contentType: "application/json", content: `{ "name":"Dasd", "age":15 }`, code: http.StatusOK},
}

var deleteRequest = []requestTest{
	{desc: "delete a user", method: "PUT", path: "/api/v1/users/61cf50b9caf825afd4391af6", contentType: "application/json", content: `{"name":"John", "age":23 }`, code: http.StatusOK},
}

func TestRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "company")
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)

	srv := httptest.NewServer(app.New(db))
	defer srv.Close()

	for _, f := range deleteRequest {
		t.Run(f.desc, func(t *testing.T) {
			if err = hitEndpoint(srv, &f); err != nil {
				t.Error(err)
			}
		})
	}
}

func hitEndpoint(srv *httptest.Server, f *requestTest) (err error) {
	req, err := http.NewRequest(f.method, srv.URL+f.path, bytes.NewBufferString(f.content))
	req.Header.Set("Content-Type", f.contentType)
	if err != nil {
		return
	}

	res, err := srv.Client().Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != f.code {
		return fmt.Errorf("expected %v; got %v", f.code, res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	if res.StatusCode == http.StatusOK {
		var v interface{}
		if err := json.Unmarshal(b, &v); err != nil {
			return fmt.Errorf("could not unmarshal body: %v", err)
		}
	}

	return
}
