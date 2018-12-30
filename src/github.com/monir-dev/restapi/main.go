package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// book Model
type Book struct {
	ID     int    `json:id`
	Title  string `json:string`
	Author string `json:author`
	Year   string `json:year`
}

// empty slice
var books []Book

func main() {
	// init router
	r := mux.NewRouter()

	// Mock Data -@todo - implement DB
	books = append(books, Book{ID: 1, Title: "Golang Pointer", Author: "Mr Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutinges", Author: "Mr Gorouting", Year: "2011"},
		Book{ID: 3, Title: "Golang router", Author: "Mr router", Year: "2012"},
		Book{ID: 4, Title: "Golang Currency", Author: "Mr Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Good", Year: "2014"})

	// Routes consist of a path and a handler function.
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":4000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var param = mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])

	for _, book := range books {
		if id == book.ID {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if id == item.ID {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
