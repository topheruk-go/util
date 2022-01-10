package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	app01 "github.com/topheruk/go/example/web/app-01"
)

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
	a := app01.New()

	log.Println(os.ExpandEnv("Listening... http://$ADDR/ping"))
	return http.ListenAndServe(os.ExpandEnv("$ADDR"), a)
}
