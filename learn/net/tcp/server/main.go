package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\v", err)
		os.Exit(1)
	}
}

func run() (err error) {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("listen error, err=%s", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept error, err=%s", err)
			return
		}
		go handleConn(conn)
		log.Printf("connection accepted %v", conn.RemoteAddr())
	}
	return
}
