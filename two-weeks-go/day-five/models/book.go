// models/book.go
package models

type Book struct { // exported — capital B
	Title  string
	Author string // unexported — lowercase
}
