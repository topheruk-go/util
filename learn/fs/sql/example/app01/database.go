package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/topheruk/go/src/encoding"
)

type Service struct {
	db *sqlx.DB
}

func NewService(ctx context.Context, dataSource string) *Service {
	db := sqlx.MustOpen("sqlite3", dataSource)
	migrate(ctx, db)
	return &Service{db: db}
}

func migrate(ctx context.Context, db *sqlx.DB) error {
	stmt := `
	DROP TABLE IF EXISTS user;
	CREATE TABLE IF NOT EXISTS user (
		"id"	BLOB,
		"email"	TEXT NOT NULL UNIQUE,
		"password" BLOB NOT NULL,
		"created_at"	DATETIME NOT NULL,
		PRIMARY KEY("id")
	);`
	db.MustExecContext(ctx, stmt)
	// TODO: will this fail at anytime?
	go db.MapperFunc(encoding.ToSnake)
	return nil
}

func (s *Service) Query(ctx context.Context, query string, args ...interface{}) error {
	return Query(ctx, s.db, query, args...)
}

func Query(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {
	stmt, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	switch len(args) {
	case 0:
		return fmt.Errorf("no arguments present")
	case 1:
		if isSlice(args[0]) {
			return SelectMany(ctx, stmt, nil, args[0])
		}
		return Exec(ctx, stmt, args[0])
	case 2:
		return Select(ctx, stmt, args[0], args[1])
	default:
		return fmt.Errorf("too many arguments present")
	}
}

func Exec(ctx context.Context, stmt *sqlx.NamedStmt, input interface{}) error {
	_, err := stmt.ExecContext(ctx, input)
	return err
}

func Select(ctx context.Context, stmt *sqlx.NamedStmt, input interface{}, target interface{}) error {
	switch reflect.Indirect(reflect.ValueOf(target)).Kind() {
	case reflect.Struct:
		return SelectOne(ctx, stmt, input, target)
	case reflect.Slice:
		return SelectMany(ctx, stmt, input, target)
		// TODO: case its a map interface
	default:
		return fmt.Errorf("illegal type")
	}
}

func SelectOne(ctx context.Context, stmt *sqlx.NamedStmt, input interface{}, target interface{}) error {
	return stmt.GetContext(ctx, target, input)
}

func SelectMany(ctx context.Context, stmt *sqlx.NamedStmt, input interface{}, target interface{}) error {
	if input == nil {
		input = struct{}{}
	}
	return stmt.SelectContext(ctx, target, input)
}

func isSlice(v interface{}) bool { return reflect.Indirect(reflect.ValueOf(v)).Kind() == reflect.Slice }
