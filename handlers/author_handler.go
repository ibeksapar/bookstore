package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var authors = []models.Author{
	{ID: 1, Name: "Alan A. A. Donovan"},
}

var authorIDCounter = 2

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	author.ID = authorIDCounter
	authorIDCounter++
	authors = append(authors, author)
	json.NewEncoder(w).Encode(author)
}
