package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "encoding/json"
    "math/rand"
    "strconv"
    // "fmt"
)

// Init books var as a slice Book struct
var books []Book

// Book Struct (Model)
type Book struct {
    ID      string  `json:"id"`
    Isbn    string  `json:"isbn"`
    Title   string  `json:"title"`
    Author  *Author `json:"author"`
}

// Author Struct (Model)

type Author struct {
    Firstname   string  `json:"firstname"`
    Lastname   string  `json:"lastname"`
}

// get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get params
    // loop through books and find with id
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000)) // MockID not safe could generate same id
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"]{
        books = append(books[:index], books[index+1:]...)
        var book Book
        _ = json.NewDecoder(r.Body).Decode(&book)
        book.ID = params["id"]
        books = append(books, book)
        json.NewEncoder(w).Encode(book)
        return
        }
    }
    json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"]{
        books = append(books[:index], books[index+1:]...)
        break
        }
    }
    json.NewEncoder(w).Encode(books)
}

func main(){
    // init router
    r := mux.NewRouter()

    // Mock Data - @todo - implement DB
    books = append(books, Book{ID: "1", Isbn:"32432", Title: "The Hobbit", Author: &Author{Firstname: "JRR", Lastname: "Tolkien"}})

    books = append(books, Book{ID: "2", Isbn:"32321", Title: "The Trial", Author: &Author{Firstname: "Franz", Lastname: "Kafka"}})

    // route handlers
    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8855", r))
}
