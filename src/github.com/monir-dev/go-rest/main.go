package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func YourHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("working..\n"))
// }

func main() {
	// init router
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/books", getBooks).method("GET")
	r.HandleFunc("/api/books/{id}", getBooks).method("GET")
	r.HandleFunc("/api/books", createBook).method("POST")
	r.HandleFunc("/api/books/{id}", updateBook).method("PUT")
	r.HandleFunc("/api/books", deleteBook).method("DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("working..\n"))
}

// Book Struct (Model)
type Book sturct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author sturct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}