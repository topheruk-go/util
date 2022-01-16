package sql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

func (db *DB) Migrate(ctx context.Context, m map[string]string) error {
	for table, schema := range m {
		db.Add(1)
		go migrate(ctx, db.s, table, schema, db.WaitGroup)
	}
	db.Wait()
	return nil
}

func migrate(ctx context.Context, db *sqlx.DB, table, schema string, wg sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, schema))
}

func (db *DB) DropAll(ctx context.Context, m map[string]string) error {
	for table := range m {
		db.Add(1)
		go drop(ctx, db.s, table, db.WaitGroup)
	}
	db.Wait()
	return nil
}

func (db *DB) Drop(ctx context.Context, table string) {
	db.s.MustExecContext(ctx, fmt.Sprintf("DELETE TABLE IF EXISTS %s", table))
}

func drop(ctx context.Context, db *sqlx.DB, table string, wg sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, fmt.Sprintf("DELETE TABLE IF EXISTS %s", table))
}
