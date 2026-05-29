package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/codyveladev/day-ten/models"
)

var (
	booksMu sync.RWMutex
	books   = []models.Book{
		{ID: 1, Title: "The Go Programming Language", Author: "Cody"},
		{ID: 2, Title: "Learn Go", Author: "Dave"},
	}
	nextID = 3
)

func listBooks(w http.ResponseWriter, r *http.Request) {
	booksMu.RLock()
	defer booksMu.RUnlock()
	author := r.URL.Query().Get("author")
	filtered := []models.Book{}
	for _, b := range books {
		if author == "" || b.Author == author {
			filtered = append(filtered, b)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listBooks(w, r)
	case http.MethodPost:
		createBook(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
