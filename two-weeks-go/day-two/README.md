# Day 2 — Data Structures

## Objectives

- Arrays vs slices and when to use each
- `append()`, `len()`, `cap()`
- Maps: declare, read, write, delete, check existence with comma-ok
- Range-based for loops
- Variadic functions

## Project

**Student grade tracker** — read student names and scores from stdin, store in a map, then print:
- Class average
- Highest scorer
- Sorted grade list

## Key concepts to understand

### Slices vs Arrays
Arrays have a fixed size (`[3]string`). Slices are dynamic and what you'll use almost always in real Go code.

```go
arr := [3]int{1, 2, 3}       // array — fixed
s   := []int{1, 2, 3}        // slice — dynamic
s    = append(s, 4)           // grows the slice
```

### Maps
```go
scores := map[string]int{}    // declare
scores["Alice"] = 95          // write
val := scores["Alice"]        // read
delete(scores, "Alice")       // delete

val, ok := scores["Alice"]    // comma-ok: ok is false if key doesn't exist
```

### Variadic functions
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)        // pass individually
sum(nums...)        // or spread a slice
```

## Done when

- [ ] Students and scores stored in a map
- [ ] Class average calculated
- [ ] Highest scorer found
- [ ] Grade list printed in sorted order
- [ ] Input handled from stdin