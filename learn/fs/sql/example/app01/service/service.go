package service

import (
	"context"

	"github.com/topheruk/go/learn/fs/sql/example/app01/model"
	"github.com/topheruk/go/src/database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB(ctx context.Context, dataSource string) *DB {
	s := sql.New(ctx, "sqlite3", dataSource)
	return &DB{DB: s}
}

func (db *DB) InsertUser(ctx context.Context, query string, dto *model.DtoUser) (*model.User, error) {
	u, err := model.New(dto)
	if err != nil {
		return nil, err
	}
	if err = db.Query(ctx, query, &u); err != nil {
		return nil, err
	}
	return u, nil
}

func (db *DB) SearchUser(ctx context.Context, query string, v interface{}) (*model.User, error) {
	var u model.User
	if err := db.Query(ctx, query, &u, v); err != nil {
		return nil, err
	}
	return &u, nil
}
