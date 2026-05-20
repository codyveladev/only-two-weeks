package main

import (
	"fmt"
)

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}

func NewBook(title string, author string) Book {
	return Book{
		Title:     title,
		Author:    author,
		Available: true,
	}
}

func (b Book) Summary() string {
	return b.Title + " by " + b.Author
}

func (b *Book) Checkout() {
	b.Available = false
}

func (b *Book) Return() {
	b.Available = true
}

func SearchByAuthor(catalog []Book, author string) []Book {
	foundBooks := []Book{}
	for _, book := range catalog {
		if author == book.Author {
			foundBooks = append(foundBooks, book)
		}
	}
	return foundBooks
}

func PrintCatalog(catalog []Book) {
	fmt.Println("== Total Catalog ==")
	for _, book := range catalog {
		var status string
		if book.Available {
			status = "Avaliable"
		} else {
			status = "Not Avaliable"
		}
		fmt.Println(book.Summary() + " " + status)
	}
}

func main() {
	catalog := []Book{}
	b1 := NewBook("The Great Gatsby", "F. Scott Fitzgerald")
	b2 := NewBook("The Martian", "Andy Weir")
	b3 := NewBook("Atomic Habits", "James Clear")
	b4 := NewBook("The Hobbit", "J.R.R. Tolkien")
	b5 := NewBook("The Girl with the Dragon Tattoo", "Stieg Larrson")
	b2.Checkout()
	b4.Checkout()
	catalog = append(catalog, b1, b2, b3, b4, b5)
	fmt.Println(catalog)
	fmt.Println(SearchByAuthor(catalog, "Andy Weir"))
	PrintCatalog(catalog)

}
