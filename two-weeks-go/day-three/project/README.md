# Day 3 — Structs & Methods

## Project: Library Catalog

A catalog of books supporting checkout, return, search by author, and a formatted availability listing.

```bash
go run main.go
```

## What I learned

### Structs
Group related fields into a named type with fixed, typed fields — like a map but enforced at compile time.
```go
type Book struct {
    Title     string `json:"title"`
    Author    string `json:"author"`
    Available bool   `json:"available"`
}
```

### Nested structs
A struct can contain another struct as a field, accessed with chained dots.
```go
type Author struct {
    FirstName string
    LastName  string
}

type Book struct {
    Author Author
}

book.Author.FirstName
```

### Constructor functions
Go has no built-in constructors. Convention is a `New` function that sets defaults and returns the struct.
```go
func NewBook(title, author string) Book {
    return Book{Title: title, Author: author, Available: true}
}
```

### Value receivers
Operates on a copy — use for read-only methods that don't need to mutate the struct.
```go
func (b Book) Summary() string {
    return b.Title + " by " + b.Author
}
```

### Pointer receivers
Operates on the original — use when the method needs to change the struct.
```go
func (b *Book) Checkout() { b.Available = false }
func (b *Book) Return()   { b.Available = true }
```

### Struct tags
Metadata on fields read by packages like `encoding/json` at runtime. Added now, used in Day 6.
```go
Title string `json:"title"`  // no space after colon
```

## Key takeaways

- If any method on a struct uses a pointer receiver, make them all pointer receivers
- Go automatically takes the address when calling a pointer receiver method on a value — no need to manually use `&`
- Mutations happen on copies unless you're working with pointers — appending a struct to a slice copies it, so checkout before appending or store pointers in the slice
- Constructor functions are convention not syntax — nothing enforces them, they're just good practice