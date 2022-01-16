package sql

import (
	"context"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

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
			return SelectMany(ctx, stmt, args[0], nil)
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

func Select(ctx context.Context, stmt *sqlx.NamedStmt, target interface{}, input interface{}) error {
	switch reflect.Indirect(reflect.ValueOf(target)).Kind() {
	case reflect.Struct:
		return SelectOne(ctx, stmt, target, input)
	case reflect.Slice:
		return SelectMany(ctx, stmt, target, input)
		// TODO: case its a map interface
	default:
		return fmt.Errorf("illegal type")
	}
}

func SelectOne(ctx context.Context, stmt *sqlx.NamedStmt, target interface{}, input interface{}) error {
	return stmt.GetContext(ctx, target, input)
}

func SelectMany(ctx context.Context, stmt *sqlx.NamedStmt, target interface{}, input interface{}) error {
	if input == nil {
		input = struct{}{}
	}
	return stmt.SelectContext(ctx, target, input)
}

func isSlice(v interface{}) bool { return reflect.Indirect(reflect.ValueOf(v)).Kind() == reflect.Slice }
