package main

// https://www.youtube.com/watch?v=SonwZ6MF5BE

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through bood and find with id
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000)) // Mock Id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if params["id"] == item.Id {
			book.Id = params["id"]
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params

	for i, item := range books {
		if params["id"] == item.Id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Upadted_at string `json:"upadted_at"`
}

var users []User

func main() {

	// connect to database
	db, err := sql.Open("mysql", "root:@/node-start")
	defer db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	// query
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		var created_at string
		var upadted_at string
		err = rows.Scan(&id, &name, &email, &password, &created_at, &upadted_at)
		checkErr(err)
		user := User{ID: id, Name: name, Email: email, Password: password, Created_at: created_at, Upadted_at: upadted_at}

		users = append(users, user)
	}
	fmt.Println(users)

	// init router
	r := mux.NewRouter()

	// Mock Data -@todo - implement DB
	books = append(books, Book{Id: "1", Isbn: "234234", Title: "Book one", Author: &Author{Firstname: "Monir", Lastname: "Hossain"}})

	books = append(books, Book{Id: "2", Isbn: "3242", Title: "Book Two", Author: &Author{Firstname: "Monir", Lastname: "Hossain"}})

	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":4000", r))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
