package sqli

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}
