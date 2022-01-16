package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

var (
	datasourceName = "./learn/fs/sql/example/app01/.sqlite3"
)

func run() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	NewService(ctx, datasourceName)

	return nil
}
