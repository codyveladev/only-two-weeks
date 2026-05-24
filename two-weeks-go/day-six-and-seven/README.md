# Day Six — Go Task CLI

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
day-six/
├── main.go               # CLI entry point, argument parsing
├── config/
│   └── constants.go      # App-wide constants (e.g. filename)
├── io/
│   └── task.go           # Load and save tasks to JSON
├── models/
│   ├── task.go           # Task struct, constructor, methods
│   └── task_list.go      # TaskList struct and operations
└── tasks.json            # Persisted task data
```

## What I Learned

### Go Fundamentals
- **Package organization** — packages are the unit of organization in Go, not files or classes. Related types and their methods live together in the same package
- **Exported vs unexported** — capitalization controls visibility; `getNextID` stays package-private, `AddTask` is public
- **No classes, just structs** — methods are attached to structs via receivers, not defined inside them

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

### Error Handling
Go returns errors as values — no exceptions. The pattern is always:
```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("context: %w", err)  // %w wraps for unwrapping later
}
```

### CLI Arguments
`os.Args` is a plain string slice — `os.Args[0]` is the program name, real args start at `os.Args[1]`:
```go
command := os.Args[1]
title := os.Args[2]
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