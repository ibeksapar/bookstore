package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books = []models.Book{
	{ID: 1, Title: "The Go Programming Language", AuthorID: 1, CategoryID: 1, Price: 45.99},
}

var bookIDCounter = 2

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	
	http.Error(w, "Book not found", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = bookIDCounter
	bookIDCounter++
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			json.NewDecoder(r.Body).Decode(&books[i])
			books[i].ID = id
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
