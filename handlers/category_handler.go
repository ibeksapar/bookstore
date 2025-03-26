package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var categories = []models.Category{
	{ID: 1, Name: "Programming"},
}

var categoryIDCounter = 2

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	category.ID = categoryIDCounter
	categoryIDCounter++
	categories = append(categories, category)
	json.NewEncoder(w).Encode(category)
}
