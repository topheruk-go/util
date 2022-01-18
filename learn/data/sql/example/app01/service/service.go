package service

import (
	"context"

	"github.com/topheruk/go/learn/data/sql/example/app01/model"
	"github.com/topheruk/go/src/database/sqli"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sqli.DB
}

func NewDB(ctx context.Context, dataSource string) *DB {
	s := sqli.MustConnect("sqlite3", dataSource)
	return &DB{DB: s}
}

func (db *DB) InsertUser(ctx context.Context, query string, dto *model.DtoUser) (*model.User, error) {
	u, err := model.New(dto)
	if err != nil {
		return nil, err
	}
	if err = db.QueryiContext(ctx, query, &u); err != nil {
		return nil, err
	}
	return u, nil
}

func (db *DB) SearchUser(ctx context.Context, query string, v interface{}) (*model.User, error) {
	var u model.User
	if err := db.QueryiContext(ctx, query, &u, v); err != nil {
		return nil, err
	}
	return &u, nil
}
