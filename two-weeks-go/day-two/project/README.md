# Day 2 — Data Structures

## Project: Student Grade Tracker

Reads student names and scores either from stdin or a file, then prints the class average, top scorers, and a sorted leaderboard.

```bash
go run main.go              # interactive stdin mode
go run main.go students.txt # read from file
```

### File format
```
Alice,95
Bob,87
Charlie,92
```

## What I learned

### Maps
Store and look up key-value pairs. Comma-ok pattern to safely check if a key exists.
```go
scores := make(map[string]int)
scores["Alice"] = 95
val, ok := scores["Alice"] // ok is false if key doesn't exist
```

### Slices & append
Dynamic arrays that grow as needed. Used a slice of names to produce a sortable list from a map.
```go
names := []string{}
for name := range scores {
    names = append(names, name)
}
```

### sort.Slice
Sort a slice by any criteria using a comparison function. Used a secondary alphabetical sort as a tiebreaker.
```go
sort.Slice(names, func(i, j int) bool {
    if scores[names[i]] == scores[names[j]] {
        return names[i] < names[j]
    }
    return scores[names[i]] > scores[names[j]]
})
```

### Variadic functions
Functions that accept any number of arguments. Useful for helpers like `sum` or `average` that operate on a variable dataset.
```go
func calculateAverage(scores ...int) float64 { ... }
```

### bufio.Scanner + file parsing
Read a file line by line, split on a delimiter, and parse into a map.
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    parts := strings.Split(scanner.Text(), ",")
}
```

### Conditional stdin vs file input
Real CLI pattern — use `os.Args` to check if a filename was passed, fall back to stdin if not.
```go
if len(os.Args) > 1 {
    studentScores, err = loadStudentsFromFile(os.Args[1])
} else {
    studentScores = readFromStdin()
}
```

## Key takeaways

- Maps have no guaranteed order — to sort, extract keys into a slice first
- `sort.Slice` is flexible enough to sort by any field with a tiebreaker
- Keeping functions pure (data in, data out) made the stdin/file switch trivial
- `strings.TrimSpace` is essential when parsing user-created files