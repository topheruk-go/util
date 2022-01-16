package client

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/learn/fs/sql/db"
	"github.com/topheruk/go/learn/fs/sql/model"
)

type Client struct {
	idb *db.ItemDatabase
}

// TODO: create dir if not exists
func New(datasourceName string) (*Client, error) {
	db, err := sqlx.Connect("sqlite3", datasourceName)
	if err != nil {
		return nil, err
	}
	c := &Client{}
	c.init(db)
	return c, nil
}

func (c *Client) init(sqlxdb *sqlx.DB) error {
	c.idb = db.NewItemDatabase(sqlxdb)
	return nil
}

func (c *Client) CreateItem(dto *model.DtoItem) (*model.Item, error) { return c.idb.CreateItem(dto) }
func (c *Client) ReadItem(id int) (*model.Item, error)               { return c.idb.ReadItem(id) }
func (c *Client) ReadAll() ([]model.Item, error)                     { return c.idb.ReadAll() }
func (c *Client) UpdateItem(i *model.Item) error                     { return c.idb.UpdateItem(i) }
func (c *Client) DeleteItem(id uuid.UUID) error                      { return c.idb.DeleteItem(id) }
