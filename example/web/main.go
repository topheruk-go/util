package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	app "github.com/topheruk/go/example/web/app-01"
)

// docker run --name mongodb -d -p 27017:27017 mongo

func init() {
	env := flag.String("env", "", "environment varaible file")
	flag.Parse()
	if err := godotenv.Load(*env); err != nil {
		log.Fatalf("[env error]\n%v", err)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s, err := app.NewService(ctx, os.ExpandEnv("mongodb://$DB_ADDR"))
	if err != nil {
		return
	}
	defer s.Close(ctx)

	a := app.NewApp(s)

	log.Println(os.ExpandEnv("Listening... http://$ADDR/ping"))
	return http.ListenAndServe(os.ExpandEnv("$ADDR"), a)
}
