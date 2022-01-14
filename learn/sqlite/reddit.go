package sqlite

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}
type PostStore interface {
	Post(id uuid.UUID) (Thread, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id uuid.UUID) error
}
type CommentStore interface {
	Comment(id uuid.UUID) (Thread, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id uuid.UUID) error
}
