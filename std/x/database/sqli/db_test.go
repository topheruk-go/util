package sqli_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func TestConnect(t *testing.T) {
	sqlx.MustConnect("sqlite3", "./test.sqlite3").PreparexContext(context.TODO(), "")
}
