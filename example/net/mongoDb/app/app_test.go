package app_test

import (
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
	username = "topheruk"
	password = "T^*G7!Pf"
	host     = "localhost"
	cliPort  = 27017
	srvPort  = 8000
	uri      = fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, cliPort)
)

type method int

const (
	Get method = iota
	Post
	Put
	Delete
)

func TestFindAllUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "company")
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)

	srv := httptest.NewServer(app.New(db))
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1/users/", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status Ok; got %v", res.Status)
	}
}

type endpointTest struct {
	Name       string
	Id         string
	StatusCode int
}

func TestFindUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, uri, "company")
	if err != nil {
		return
	}
	defer db.Disconnect(ctx)

	srv := httptest.NewServer(app.New(db))
	defer srv.Close()

	ep := []endpointTest{
		{Name: "Finding User; OK", Id: "61ceea9b0fa3d8b8c5bd9292", StatusCode: http.StatusOK},
		{Name: "Finding User; OK", Id: "61ce33e3928e6155964a629f", StatusCode: http.StatusOK},
		{Name: "Finding User; DB Search Error", Id: "34ce33e3928e6155964a629f", StatusCode: http.StatusInternalServerError},
		{Name: "Finding User; MongoDB ObjectID convertion", Id: "34ce33e5964a629f", StatusCode: http.StatusInternalServerError},
	}

	for _, e := range ep {
		if err := getUser(srv, e); err != nil {
			t.Fatal(err)
		}
	}
}

func getUser(srv *httptest.Server, e endpointTest) error {
	res, err := http.Get(fmt.Sprintf("%s/api/v1/users/%s", srv.URL, e.Id))
	if err != nil {
		return fmt.Errorf("could not send GET request: %v", err)
	}

	if res.StatusCode != e.StatusCode {
		return fmt.Errorf("expected %v; got %v", e.StatusCode, res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	// If there is an error, there is no way to do further checks
	if res.StatusCode == http.StatusOK {
		var u database.User
		if err := json.Unmarshal(b, &u); err != nil {
			return fmt.Errorf("could not unmarshal body: %v", err)
		}
	}

	return nil
}
