package main

import (
	"fmt"
	"net/http"

	"github.com/codyveladev/day-ten/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	})
	http.HandleFunc("/books", handlers.HandleBooks)
	fmt.Println("server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
