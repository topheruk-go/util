package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/alexmullins/zip"
)

func main() {
	contents := []byte("Hello World")
	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	w, err := zipw.Encrypt(`test.txt`, `golang`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
}
