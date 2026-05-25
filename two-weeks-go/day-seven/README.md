# Day Six & Seven — Go Task CLI

A command-line task manager built in Go as part of a two-week Go learning challenge. Tasks are persisted to a local JSON file and support adding, completing, deleting, and listing.

## Usage

```bash
go run . add "Buy groceries"      # add a task
go run . list                     # list all tasks
go run . done 1                   # mark task 1 as complete
go run . delete 1                 # delete task 1
```

## Project Structure

```
day-six-and-seven/
├── main.go               # CLI entry point, argument parsing
├── config/
│   └── constants.go      # App-wide constants (e.g. filename)
├── io/
│   └── task.go           # Load and save tasks to JSON
├── models/
│   ├── task.go           # Task struct, constructor, methods
│   ├── task_test.go      # Tests for Task methods
│   ├── task_list.go      # TaskList struct and operations
│   └── task_list_test.go # Tests for TaskList operations
└── tasks.json            # Persisted task data
```

---

## Day 6 — File I/O & JSON

### Package organization
Packages are the unit of organization in Go, not files or classes. Related types and their methods live together in the same package.

### Exported vs unexported
Capitalization controls visibility — `getNextID` stays package-private, `AddTask` is public.

### Constructors
Go has no `new` keyword for custom initialization. The convention is a `NewX` function:
```go
func NewTask(id int, title string) Task { ... }
```

### Pointer vs Value Receivers
- A value receiver gets a **copy** — safe for reads, useless for mutation
- A pointer receiver gets the **real thing** — required when you need to change state
```go
func (t *Task) Complete() {
    t.Completed = true  // modifies the original
}
```

### Range Loop Copies
`for _, task := range tasks` gives you a **copy** of each element. To mutate or delete, you need the index:
```go
// ❌ modifies a copy, original unchanged
for _, task := range taskList.Tasks {
    task.Complete()
}

// ✅ modifies the real element
for i := range taskList.Tasks {
    taskList.Tasks[i].Complete()
}
```

### Deleting from a Slice
Go has no built-in slice delete. The idiomatic approach uses `append` to stitch around the element:
```go
tasks = append(tasks[:i], tasks[i+1:]...)
```

### ID Generation
Length-based IDs (`len(tasks) + 1`) break after any deletion. The correct approach is to derive the next ID from the current maximum:
```go
func getNextID(taskList TaskList) int {
    max := 0
    for _, t := range taskList.Tasks {
        if t.ID > max {
            max = t.ID
        }
    }
    return max + 1
}
```

---

## Day 7 — Testing

### The testing package
Go has testing built in — no external library needed. Test files end in `_test.go` and live next to the files they test. Test functions start with `Test` and take `*testing.T`.
```go
func TestNewTask(t *testing.T) {
    task := NewTask(1, "Buy Milk")
    if task.ID != 1 {
        t.Errorf("expected ID 1, got %d", task.ID)
    }
}
```

### Table-driven tests
The idiomatic Go pattern — define a slice of test cases and loop over them. Adding a new case is one struct literal, failures are identified by name.
```go
tests := []struct {
    name     string
    input    int
    expected int
}{
    {"empty list", 0, 1},
    {"one task",   1, 2},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        // assert
    })
}
```

### t.Run subtests
Creates a named subtest so failures report the exact case: `TestGetNextId/empty_list_returns_1` instead of just `TestGetNextId failed`.

### t.Errorf vs t.Fatalf
- `t.Errorf` — marks test failed, keeps running. Use when you want to catch multiple failures.
- `t.Fatalf` — marks test failed, stops immediately. Use when later assertions depend on earlier ones.

### Testing for errors
The `wantErr bool` pattern checks both directions — error when expected, no error when not:
```go
if (err != nil) != tt.wantErr {
    t.Errorf("got err=%v, wantErr=%v", err, tt.wantErr)
}
```

### Same-package testing
Test files in the same package can access unexported identifiers — `getNextId` stays private but is still testable without making it public.

### Benchmarks
```go
func BenchmarkGetNextId(b *testing.B) {
    // setup
    b.ResetTimer() // don't count setup time
    for i := 0; i < b.N; i++ {
        getNextId(taskList) // b.N set automatically by Go
    }
}
```

### Running tests
```bash
go test ./...                          # run all tests
go test ./... -v                       # verbose output
go test ./... -cover                   # show coverage %
go test ./... -bench=. -benchmem       # run benchmarks
go test ./models/... -coverprofile=coverage.out
go tool cover -html=coverage.out       # open coverage in browser
```

## Key takeaways

- Test files live next to the code they test — same package, `_test.go` suffix
- Table-driven tests are idiomatic Go — write the structure once, add cases as data
- `(err != nil) != tt.wantErr` catches both unexpected errors and missing errors in one line
- Same-package tests can access unexported functions — no need to make things public just for testing
- Benchmarks are for hot paths — `getNextId` on a task list will never need optimization in practice
- Coverage measures lines executed, not correctness — 80% is a signal, not a goal