package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books = []models.Book{
	{ID: 1, Title: "The Alchemist", AuthorID: 1, CategoryID: 1, Price: 12.99},
	{ID: 2, Title: "To Kill a Mockingbird", AuthorID: 2, CategoryID: 2, Price: 10.50},
	{ID: 3, Title: "1984", AuthorID: 3, CategoryID: 3, Price: 8.99},
	{ID: 4, Title: "The Great Gatsby", AuthorID: 4, CategoryID: 2, Price: 11.25},
	{ID: 5, Title: "The Little Prince", AuthorID: 5, CategoryID: 1, Price: 9.40},
	{ID: 6, Title: "Sapiens: A Brief History of Humankind", AuthorID: 6, CategoryID: 4, Price: 14.80},
}


var bookIDCounter = 2

func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	categoryParam := query.Get("category")
	pageParam := query.Get("page")
	limitParam := query.Get("limit")

	// filter
	filteredBooks := books
	if categoryParam != "" {
		categoryID, err := strconv.Atoi(categoryParam)

		if err == nil {
			var temp []models.Book
			for _, book := range books {
				if book.CategoryID == categoryID {
					temp = append(temp, book)
				}
			}

			filteredBooks = temp
		}
	}

	// pagination
	page := 1
	limit := len(filteredBooks)

	if p, err := strconv.Atoi(pageParam); err == nil {
		page = p
	}

	if l, err := strconv.Atoi(limitParam); err == nil {
		limit = l
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filteredBooks) {
		start = len(filteredBooks)
	}

	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	json.NewEncoder(w).Encode(filteredBooks[start:end])
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

	// validation
	if book.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if book.Price <= 0 {
		http.Error(w, "Price must be greater than 0", http.StatusBadRequest)
		return
	}

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
