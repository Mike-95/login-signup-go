package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/signup", signupHandler)
	mux.HandleFunc("/login", loginHandler)

	log.Println("Starting server on :4040")
	err := http.ListenAndServe(":4040", mux)

	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("home"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup page..."))

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page.."))

}
