package main

import (
	"net/http"

	"github.com/codyveladev/day-eleven/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/books", handlers.ListBooks)
	r.Post("/books", handlers.CreateBook)
	// TODO
	//r.Get("/books/{id}", handlers.GetBook)
	//r.Put("/books/{id}", handlers.UpdateBook)
	//r.Delete("/books/{id}", handlers.DeleteBook)

	http.ListenAndServe(":8080", r)
}
