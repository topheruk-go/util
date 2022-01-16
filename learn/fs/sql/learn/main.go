package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/topheruk/go/learn/fs/sql/client"
	"github.com/topheruk/go/learn/fs/sql/model"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cli, err := client.New("./learn/fs/sql/.sqlite3")
	if err != nil {
		return err
	}

	resItem := &model.DtoItem{Title: "Bar", CreatedAt: time.Now()}
	newItem, err := cli.CreateItem(resItem)
	if err != nil {
		return err
	}

	fmt.Println(newItem)

	// read all
	items, err := cli.ReadAll()
	if err != nil {
		return err
	}

	for _, i := range items {
		fmt.Println(i)
	}

	return nil
}
