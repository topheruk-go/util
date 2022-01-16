package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateUser(ctx context.Context, db *sqlx.DB, dto *DtoUser) (*User, error) {
	stmt := `
	INSERT INTO users VALUES (?, ?, ?, ?)`
	user, err := NewUser(dto)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user: %w", err)
	}
	if _, err = db.ExecContext(ctx, stmt,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to add new user to database: %w", err)
	}
	return user, nil
}

func ReadUserByEmail(ctx context.Context, db *sqlx.DB, email string) (*User, error) {
	stmt := `
	SELECT * FROM users WHERE email = ?`
	u := &User{}
	// change this to named exec?
	if err := db.GetContext(ctx, u, stmt, email); err != nil {
		return nil, fmt.Errorf("invalid email entered: %w", err)
	}
	return u, nil
}

func ReadUserByID(ctx context.Context, db *sqlx.DB, id uuid.UUID) (*User, error) {
	stmt := `
	SELECT * FROM users WHERE id = ?`
	u := &User{}
	if err := db.GetContext(ctx, u, stmt, id); err != nil {
		return nil, fmt.Errorf("invalid email entered: %w", err)
	}
	return u, nil
}

func ReadUser(ctx context.Context, db *sqlx.DB, v interface{}) {}
