package pg

import (
	"fmt"
	"os"
)

// Environment Variables
//
// POSTGRES_USER=postgres
// POSTGRES_PASSWORD=postgres
// POSTGRES_DB=postgres
// POSTGRES_HOSTNAME=localhost
// DB_PORT=5432
// SSL_MODE=disable
var DefaultConnString = ConnString{"POSTGRES_HOSTNAME", "DB_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB", "SSL_MODE"}

type ConnString struct {
	hostname string
	port     string
	username string
	password string
	dbName   string
	sslMode  string
}

func NewConnString(hostname, port, username, password, dbName, sslMode string) *ConnString {
	return &ConnString{hostname, port, username, password, dbName, sslMode}
}

func (c *ConnString) String() string {
	return fmt.Sprintf("host=${%s} port=${%s} user=${%s} password=${%s} dbname=${%s} sslmode=${%s}", c.hostname, c.port, c.username, c.password, c.dbName, c.sslMode)
}

func (c *ConnString) ExpandEnv() string {
	return os.ExpandEnv(c.String())
}
