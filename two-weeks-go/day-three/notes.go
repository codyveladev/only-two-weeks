package main

import (
	"fmt"
)

type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title     string `json:"title"`
	Author    Author `json:"author"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}

func NewBook(title string, author Author, year int) Book {
	return Book{
		Title:     title,
		Author:    author,
		Year:      year,
		Available: true,
	}
}

// Value reciever - works on a COPY
func (b Book) Summary() string {
	return b.Title + " by " + b.Author.FirstName + " " + b.Author.LastName
}

// pointer receiver - works on the original
func (b *Book) Checkout() {
	b.Available = false
}

// Changing instance so needs pointer reciever
func (b *Book) Return() {
	b.Available = true
}

func main() {
	b1 := NewBook("Moby Dick", Author{FirstName: "Cody", LastName: "Vela"}, 2020)
	b2 := NewBook("Cody's Journal", Author{FirstName: "Cody", LastName: "Vela"}, 2019)
	b3 := NewBook("Frankenstein", Author{FirstName: "Cody", LastName: "Vela"}, 1992)
	b3ptr := &b3
	b3.Checkout()
	fmt.Println(b1, b2, *b3ptr)
	b3ptr.Return()
	fmt.Println(b1, b2, *b3ptr)

	fmt.Println("Address: ", &b3ptr, "Value: ", *b3ptr)
}
