package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/codyveladev/day-eleven/models"
	"github.com/go-chi/chi/v5"
)

var (
	booksMu sync.RWMutex
	books   = []models.Book{
		{ID: 1, Title: "The Go Programming Language", Author: "Cody"},
		{ID: 2, Title: "Learn Go", Author: "Dave"},
	}
	nextID = 3
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
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

func GetBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	booksMu.RLock()
	defer booksMu.RUnlock()
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "book not found", http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()
	for i, book := range books {
		if book.ID == id {
			books[i] = models.Book{ID: id, Title: newBook.Title, Author: newBook.Author}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	booksMu.Lock()
	defer booksMu.Unlock()
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}
