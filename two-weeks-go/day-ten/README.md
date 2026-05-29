# Day 10 — net/http

## Project: Books REST API (part 1 of 3)

A JSON REST API for managing a books resource. In-memory storage protected by a mutex. Days 11–12 will add proper routing/middleware and swap to SQLite.

```bash
go run main.go
# server starting on :8080
```

## Project Structure

```
day-ten/
├── go.mod
├── main.go               # server setup + route registration
├── handlers/
│   └── books.go          # HandleBooks dispatcher + list/create
└── models/
    └── book.go           # Book struct + json tags
```

### main.go
Registers two handlers — a root "hello world" at `/` and `/books` which dispatches based on HTTP method. Starts the server on port 8080.

### handlers/books.go
Holds the in-memory book storage (a package-level slice protected by a mutex) and three handler functions:
- `HandleBooks` — method dispatcher, branches on `r.Method`
- `listBooks` / `listBooksByAuthor` — GET handlers
- `createBook` — POST handler

### models/book.go
The Book struct with json tags. Lives in its own package so future code (database, tests) can import it without depending on handlers.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/books` | List all books |
| GET | `/books?author=Name` | Filter by author |
| POST | `/books` | Create a new book |

## Testing with curl

```bash
# list all books
curl http://localhost:8080/books

# filter by author
curl "http://localhost:8080/books?author=Cody"

# create a new book (use curl.exe on Windows PowerShell)
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Learning Go","author":"Bodner"}' \
  http://localhost:8080/books

# method not allowed returns 405
curl -X DELETE http://localhost:8080/books
```

## What I learned

### The handler signature
Every HTTP handler in Go has the same signature. This never changes — middleware, routes, frameworks all build on this.
```go
func(w http.ResponseWriter, r *http.Request)
```
- `w` — write the response into it (headers, status, body)
- `r` — read everything about the incoming request

### http.HandleFunc and ListenAndServe
The simplest possible server registers handlers with the default ServeMux and starts listening.
```go
http.HandleFunc("/books", handlers.HandleBooks)
http.ListenAndServe(":8080", nil)
```
The `nil` tells Go to use the default mux. Tomorrow's router upgrade replaces this.

### Returning JSON
The standard pattern is: set the Content-Type header, optionally write a status code, then encode the value into the response writer.
```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(book)
```

### Reading JSON bodies
Symmetric to encoding — decode directly from the request body into a struct. Handle decode errors with a 400.
```go
var book models.Book
if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
    http.Error(w, "invalid JSON", http.StatusBadRequest)
    return
}
```

### Method routing via switch
Before middleware/routing frameworks, the way to handle multiple methods on the same path is a switch on `r.Method`.
```go
switch r.Method {
case http.MethodGet:    listBooks(w, r)
case http.MethodPost:   createBook(w, r)
default: http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
```

### Query parameters
Use `r.URL.Query()` to read query strings. `Get()` returns an empty string if the key isn't present — convenient for "no filter applied" defaults.
```go
author := r.URL.Query().Get("author")  // /books?author=Bodner
```

### Header ordering
Headers must be set BEFORE writing the body. Once you call `Encode`, `Write`, or `Fprintln` on the ResponseWriter, the headers are sent and locked in. Any later `Header().Set()` calls are silently ignored.

### http.Error returns text/plain
`http.Error` writes a status code and a plain-text body. It does NOT return — the handler keeps running. Always `return` after calling it.
```go
http.Error(w, "invalid JSON", http.StatusBadRequest)
return  // critical
```

### HTTP status codes
- `200 OK` — default
- `201 Created` — successful POST creating a resource
- `400 Bad Request` — malformed input
- `404 Not Found` — resource doesn't exist
- `405 Method Not Allowed` — method not supported at this path
- `500 Internal Server Error` — server-side failure

### Concurrency in handlers
`net/http` serves requests concurrently — each request runs in its own goroutine. The package-level `books` slice is shared state, so reads and writes must be protected by a mutex (everything from Day 9 applies).
```go
var booksMu sync.Mutex
booksMu.Lock()
defer booksMu.Unlock()
// access books
```

## Key takeaways

- Every Go handler has the same signature — `func(http.ResponseWriter, *http.Request)`
- Set headers before writing the body — order matters
- Always `return` after `http.Error` or you'll write two responses
- `net/http` runs each request in its own goroutine — shared state needs locking
- `json.NewEncoder(w).Encode(v)` is the standard pattern for writing JSON responses
- `json.NewDecoder(r.Body).Decode(&v)` is the symmetric pattern for reading JSON bodies
- The default ServeMux can't do path parameters cleanly — that's tomorrow's problem to solve with a real router
- Status codes matter — 201 for POST success, 400 for bad input, 405 for wrong method