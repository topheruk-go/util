package sql

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gotest.tools/v3/assert"
)

var tblPerson = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

var tblPlace = `CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Place struct {
	Country string
	City    *string
	TelCode int
}

func TestSelectSQL(t *testing.T) {
	db := sqlx.MustOpen("sqlite3", ":memory:")
	defer db.Close()

	// Create tables
	db.MustExec(tblPerson)
	db.MustExec(tblPlace)

	// Insert data
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	tx.Commit()

	// Basic query
	var ps []Person
	err := db.Select(&ps, "SELECT * FROM person ORDER BY first_name ASC")
	assert.NilError(t, err)
	// assert.Equal(t, len(ps), 2)

	var pp []Person
	assert.NilError(t, Query(db, &pp, "SELECT * FROM person ORDER BY first_name ASC"))
	assert.Equal(t, len(ps), len(pp))

	var p Person
	assert.NilError(t, Query(db, &p, "SELEct * FROM person WHERE first_name = $1", "John"))
	assert.Equal(t, p.FirstName, "John")

	// error
	err = Query(db, Person{}, "SELEct * FROM person WHERE first_name = $1", "John")
	assert.ErrorContains(t, err, "must pass by reference")

	err = Query(db, []Person{}, "SELECT * FROM person ORDER BY first_name ASC")
	assert.ErrorContains(t, err, "must pass by reference")
}
