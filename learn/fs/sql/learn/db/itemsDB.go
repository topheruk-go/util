package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/learn/fs/sql/model"
)

type ItemDatabase struct {
	*sqlx.DB
}

func NewItemDatabase(db *sqlx.DB) *ItemDatabase {
	db.MustExec(migrate[CreateItems])
	return &ItemDatabase{db}
}

func (idb *ItemDatabase) ReadItem(id int) (*model.Item, error) {
	i := &model.Item{}
	if err := idb.Get(i, stmts[ReadItem], id); err != nil {
		return nil, fmt.Errorf("error reading item: %w", err)
	}
	return i, nil
}

func (idb *ItemDatabase) ReadAll() ([]model.Item, error) {
	ii := []model.Item{}
	if err := idb.Select(&ii, stmts[ReadAllItems]); err != nil {
		return nil, fmt.Errorf("error reading all items: %w", err)
	}
	return ii, nil
}

func (idb *ItemDatabase) CreateItem(dto *model.DtoItem) (*model.Item, error) {
	i, err := model.NewItem(dto)
	if err != nil {
		return nil, err
	}
	if _, err := idb.Exec(stmts[InsertItem], i.ID, i.Title, i.CreatedAt); err != nil {
		return nil, fmt.Errorf("error creating item: %w", err)
	}
	return i, nil
}

func (idb *ItemDatabase) UpdateItem(i *model.Item) error {
	if err := idb.Get(i, stmts[UpdateItem], i.Title, i.CreatedAt, i.ID); err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}
	return nil
}

func (idb *ItemDatabase) DeleteItem(id uuid.UUID) error {
	if _, err := idb.Exec(stmts[DeleteItem], id); err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}
	return nil
}
