package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type testcase struct {
	pathname    string
	method      string
	contentType string
	content     string
	status      int
}

func TestPing(t *testing.T) {
	ser := New(sqlx.MustConnect("sqlite3", ":memory:"))
	srv := httptest.NewServer(ser)

	defer func() { ser.Close(); srv.Close() }()

	var tt = []testcase{
		{pathname: "/ping", status: http.StatusOK},
		// TODO: trailing slashes
		// {pathname: "/ping/", method: http.MethodGet, status: http.StatusOK},

		{pathname: "/err", status: http.StatusNotFound},

		{pathname: "/person", method: http.MethodPost, content: `{"name":"John"}`, status: http.StatusCreated},
		{pathname: "/person", method: http.MethodPost, content: `{"name":30}`, status: http.StatusBadRequest},
		{pathname: "/person", method: http.MethodPost, content: `{"name:30}`, status: http.StatusBadRequest},
		{pathname: "/person", method: http.MethodPost, content: `{"name":"John"}`, status: http.StatusInternalServerError},
		// TODO: expect: BadRequest
		{pathname: "/person", method: http.MethodPost, content: `{}`, status: http.StatusInternalServerError},
		{pathname: "/person", method: http.MethodPost, content: `{"name":"Mary"}`, status: http.StatusCreated},

		{pathname: "/person"},
		{pathname: "/person?id=1&name=John"},

		{pathname: "/person/1"},
		{pathname: "/person/one", status: http.StatusBadRequest},
		{pathname: "/person/2"},
		{pathname: "/person/3", status: http.StatusNotFound},

		{pathname: "/person/1", method: http.MethodDelete, status: http.StatusNoContent},
		{pathname: "/person/3", method: http.MethodDelete, status: http.StatusInternalServerError},

		// {pathname: "/person/2", method: http.MethodPut}
	}

	for i, tc := range tt {
		if tc.contentType == "" {
			tc.contentType = "application/json"
		}
		if tc.status == 0 {
			tc.status = http.StatusOK
		}

		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
			req, err := http.NewRequest(tc.method, srv.URL+tc.pathname, strings.NewReader(tc.content))
			assert.Equal(t, err, nil)

			// TODO: do I need this?
			req.Header.Add("Content-Type", tc.contentType)

			res, err := srv.Client().Do(req)
			assert.Equal(t, err, nil)

			assert.Equal(t, res.StatusCode, tc.status)
		})
	}
}
