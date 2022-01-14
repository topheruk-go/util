package example

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/topheruk/go/learn/sqlite"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (t sqlite.Thread, err error) {
	return t, fmt.Errorf("error getting thread: %w",
		s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id))
}

func (s *ThreadStore) Threads() (tt []sqlite.Thread, err error) {
	return tt, fmt.Errorf("%w",
		s.Select(&tt, `SELECT * FROM threads`))
}

func (s *ThreadStore) CreateThread(t *sqlite.Thread) error {
	return fmt.Errorf("%w",
		s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
			t.ID,
			t.Title,
			t.Description))
}

func (s *ThreadStore) UpdateThread(t *sqlite.Thread) error {
	return fmt.Errorf("error creating thread: %w",
		s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3) RETURNING *`,
			t.Title,
			t.Description,
			t.ID))
}

func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
