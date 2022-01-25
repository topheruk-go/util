package sqli

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type namedExecer interface {
	ExecContext(ctx context.Context, arg interface{}) (sql.Result, error)
}

type execer interface {
	ExecContext(ctx context.Context, arg ...interface{}) (sql.Result, error)
}

type preparer interface {
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

func preparedQuery(ctx context.Context, p preparer, query string, args ...interface{}) error {
	stmt, err := p.PreparexContext(ctx, query)
	if err != nil {
		return err
	}
	switch len(args) {
	case 0:
		return exec(ctx, stmt, nil)
	case 1:
		if isSlice(args[0]) {
			return namedSelect(ctx, stmt, args[0], nil)
		}
		return namedExec(ctx, stmt, args[0])
	case 2:
		return namedSelect(ctx, stmt, args[0], args[1])
	default:
		return fmt.Errorf("too many arguments present")
	}
}

func preparedNamedQuery(ctx context.Context, p preparer, query string, args ...interface{}) error {
	stmt, err := p.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	switch len(args) {
	case 0:
		return namedExec(ctx, stmt, nil)
	case 1:
		if isSlice(args[0]) {
			return namedSelect(ctx, stmt, args[0], nil)
		}
		return namedExec(ctx, stmt, args[0])
	case 2:
		return namedSelect(ctx, stmt, args[0], args[1])
	default:
		return fmt.Errorf("too many arguments present")
	}
}

// stmt.GetContext() GetContext(ctx context.Context, dest interface{}, args ...interface{}) error
// stmt.SelectContext() SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error

func namedExec(ctx context.Context, e namedExecer, input interface{}) error {
	if input == nil {
		input = struct{}{}
	}
	_, err := e.ExecContext(ctx, input)
	return err
}

func namedSelect(ctx context.Context, e *sqlx.NamedStmt, target interface{}, input interface{}) error {
	switch reflect.Indirect(reflect.ValueOf(target)).Kind() {
	case reflect.Struct:
		return e.GetContext(ctx, target, input)
		// TODO: case its a map interface
	case reflect.Slice, reflect.Map:
		if input == nil {
			input = struct{}{}
		}
		return e.SelectContext(ctx, target, input)
	default:
		return fmt.Errorf("illegal type")
	}
}

func isSlice(v interface{}) bool { return reflect.Indirect(reflect.ValueOf(v)).Kind() == reflect.Slice }
