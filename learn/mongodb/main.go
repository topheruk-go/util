package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

// no auth
// docker run --name mongodb -d -p 27017:27017 mongo

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s, err := NewService(ctx, "mongodb://localhost:27017")
	if err != nil {
		return
	}
	defer s.Disconnect(ctx)

	dbNames, _ := s.ListDatabaseNames(ctx, s.Nil())
	for _, db := range dbNames {
		fmt.Printf("%v\n", db)
	}

	// if err = doSomething(ctx, s); err != nil {
	if err = doSomething2(ctx, s); err != nil {
		return
	}

	return
}

func doSomething(ctx context.Context, s *Service) (err error) {
	// INSERT ONE
	dto := &DTO{Email: "mike@gmail.com", Password: "passw0rd"}
	user, err := insertOne(ctx, s, dto)
	if err != nil {
		return
	}
	fmt.Printf("added: %v\n", user)

	// FIND ONE
	var found User
	if err = findOne(ctx, s, user.ID.Hex(), &found); err != nil {
		return
	}

	fmt.Printf("found: %v\n", found)

	// DELETE ONE
	if err = deleteOne(ctx, s, user.ID.Hex()); err != nil {
		return
	}
	fmt.Println("delete successful")

	return nil
}

func doSomething2(ctx context.Context, s *Service) (err error) {
	// INSERT ONE
	dto := &DTO{Email: "mike@gmail.com", Password: "passw0rd"}
	user, err := insertOne(ctx, s, dto)
	if err != nil {
		return
	}

	fmt.Printf("added: %v with password %v\n", user.Email, user.PasswordHash)

	// SHOW ALL
	users, err := s.FindUserMany(ctx, s.Nil())
	if err != nil {
		return
	}

	fmt.Printf("there are %d in the collection\n", len(users))

	// DELETE ALL
	if err = s.DeleteUserAll(ctx); err != nil {
		return
	}

	fmt.Println("deleted all")

	// SHOW ALL
	users, err = s.FindUserMany(ctx, s.Nil())
	if err != nil {
		return
	}

	fmt.Printf("there are %d in the collection\n", len(users))

	for i, u := range users {
		fmt.Printf("%d - %s\n", i, u)
	}

	return nil
}

func insertOne(ctx context.Context, s *Service, dto *DTO) (*User, error) {
	user, err := dto.New()
	if err != nil {
		return nil, err
	}

	err = s.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func findOne(ctx context.Context, s *Service, id string, found *User) error {
	oid, err := s.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.FindUser(ctx, s.FilterByID(oid), found)
}

func deleteOne(ctx context.Context, s *Service, id string) error {
	oid, err := s.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.DeleteUser(ctx, s.FilterByID(oid))
}

// ObjectID("61ddf3d4eba41b87dc8d5920") delete
