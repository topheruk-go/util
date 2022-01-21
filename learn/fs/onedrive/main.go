package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

const (
	path = "C:/Users/kraff/OneDrive - Fashion Retail Academy/loans/master/"
)

func run() error {
	ff, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for i, f := range ff {
		if err := os.Rename(path+f.Name(), path+fmt.Sprintf("%d.pdf", i)); err != nil {
			return err
		}
	}
	return nil
}

// outstanding_list

// extension
// "id" uuid primary key
// "loan_id" foreign key reference
//
