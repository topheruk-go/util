package sqli_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestConnect(t *testing.T) {
	sqlx.MustConnect("", "").PreparexContext(context.TODO(), "")
}
