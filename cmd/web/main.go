package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type application struct {
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
