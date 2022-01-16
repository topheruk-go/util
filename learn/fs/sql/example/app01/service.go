package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/src/database/sql"
	"github.com/topheruk/go/src/encoding"

	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	sync.WaitGroup
	db *sqlx.DB
}

func newService(ctx context.Context, dataSource string) *service {
	db := sqlx.MustOpen("sqlite3", dataSource)
	return &service{db: db}
}

func (s *service) Migrate(ctx context.Context, m map[string]string) error {
	for table, schema := range m {
		s.Add(1)
		go migrate(ctx, s.db, table, schema, s.WaitGroup)
	}
	s.Wait()
	s.db.MapperFunc(encoding.ToSnake)
	return nil
}

func migrate(ctx context.Context, db *sqlx.DB, table, schema string, wg sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, schema))
}

func (s *service) DropAll(ctx context.Context, m map[string]string) error {
	for table := range m {
		s.Add(1)
		go drop(ctx, s.db, table, s.WaitGroup)
	}
	s.Wait()
	return nil
}

func (s *service) Drop(ctx context.Context, table string) {
	s.db.MustExecContext(ctx, fmt.Sprintf("DELETE TABLE IF EXISTS %s", table))
}

func drop(ctx context.Context, db *sqlx.DB, table string, wg sync.WaitGroup) {
	defer wg.Done()
	db.MustExecContext(ctx, fmt.Sprintf("DELETE TABLE IF EXISTS %s", table))
}

func (s *service) InsertUser(ctx context.Context, query string, dto *DtoUser) (*User, error) {
	u, err := NewUser(dto)
	if err != nil {
		return nil, err
	}

	if err = s.Query(ctx, query, &u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) SearchUser(ctx context.Context, query string, v interface{}) (*User, error) {
	var u User
	if err := s.Query(ctx, query, &u, v); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *service) Query(ctx context.Context, query string, args ...interface{}) error {
	return sql.Query(ctx, s.db, query, args...)
}
