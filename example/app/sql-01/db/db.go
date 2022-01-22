package db

import "github.com/topheruk/go/example/app/sql-01/model"

type DB struct {
	Forms []model.LoanForm
}

func New() *DB {
	db := &DB{
		Forms: []model.LoanForm{},
	}
	return db
}
