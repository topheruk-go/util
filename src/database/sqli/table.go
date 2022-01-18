package sqli

import (
	"context"
	"sync"

	"github.com/jmoiron/sqlx"
)

func (db *DB) Migrate(queries ...string) error {
	return db.MigrateContext(context.Background(), queries...)
}

func (db *DB) MigrateContext(ctx context.Context, queries ...string) error {
	var wg sync.WaitGroup
	for _, schema := range queries {
		wg.Add(1)
		go migrate(ctx, db.DB, schema, &wg)
	}
	wg.Wait()
	return nil
}

func migrate(ctx context.Context, db *sqlx.DB, schema string, wg *sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, schema)
}

func (db *DB) DropAll(ctx context.Context, qs []string) error {
	var wg sync.WaitGroup
	for _, query := range qs {
		wg.Add(1)
		go drop(ctx, db.DB, query, &wg)
	}
	wg.Wait()
	return nil
}

func (db *DB) Drop(ctx context.Context, query string) {
	db.MustExecContext(ctx, query)
}

func drop(ctx context.Context, db *sqlx.DB, query string, wg *sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, query)
}
