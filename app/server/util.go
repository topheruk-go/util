package main

import (
	"os"
	"path/filepath"
)

func btof(b []byte) (*os.File, error) {
	f, err := os.CreateTemp("./app/tmp", "*.pdf")
	if err != nil {
		return nil, err
	}
	_, err = f.Write(b)
	return f, err
}

func removeAll(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
