package db

import "context"

type Database interface {
	Insert(ctx context.Context, v interface{}) error
	Search(ctx context.Context, v interface{}) error
}
