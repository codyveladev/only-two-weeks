package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codyveladev/day-eleven/handlers"
	"github.com/codyveladev/day-eleven/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Common Middleware
	r.Use(middleware.LoggingMiddleware)

	//Public Routes
	r.Get("/books", handlers.ListBooks)
	r.Get("/books/{id}", handlers.GetBook)

	//Protected Routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireApiKey)
		r.Post("/books", handlers.CreateBook)
		r.Put("/books/{id}", handlers.UpdateBook)
		r.Delete("/books/{id}", handlers.DeleteBook)
	})

	server := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Println("server started on :8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("forced shutdown:", err)
	}
	log.Println("server stopped")
}
