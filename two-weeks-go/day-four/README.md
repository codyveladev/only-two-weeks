# Day 4 — Interfaces & Errors

## Project: Shape Area Calculator

A CLI tool that calculates area and perimeter for circles, rectangles, and triangles with full error handling for invalid dimensions.

```bash
go run main.go
```

## What I learned

### Interfaces
Define a set of methods a type must have. Any type that implements those methods satisfies the interface automatically — no explicit declaration needed (implicit implementation).
```go
type Shape interface {
    Area()      float64
    Perimeter() float64
}
```

### The error interface
`error` is just a built-in interface with one method. Any type that has `Error() string` satisfies it.
```go
type error interface {
    Error() string
}
```

### Custom error types
Structs that satisfy the `error` interface. Useful when you need to attach structured data to an error.
```go
type InvalidDimensionError struct {
    Shape string
    Value float64
}

func (e InvalidDimensionError) Error() string {
    return fmt.Sprintf("%s cannot have negative dimension: %.2f", e.Shape, e.Value)
}
```

### Sentinel errors
Package-level error variables that represent a specific known error. Used with `errors.Is`.
```go
var ErrNegativeDimension = errors.New("negative dimension")
```

### Error wrapping with %w
Wraps one error inside another so the original is still inspectable by `errors.Is` and `errors.As`.
```go
return Circle{}, fmt.Errorf("circle: %w %w", ErrNegativeDimension, InvalidDimensionError{...})
```

### errors.Is vs errors.As
```go
// Is — checks if a specific sentinel is anywhere in the error chain
errors.Is(err, ErrNegativeDimension)

// As — extracts a specific error type from the chain so you can read its fields
var dimErr InvalidDimensionError
errors.As(err, &dimErr) // dimErr.Shape and dimErr.Value now accessible
```

### Type assertions & type switches
Extract the concrete type from an interface variable. Type switches are cleaner when checking multiple types.
```go
switch v := s.(type) {
case Circle:
    fmt.Println(v.Radius) // v is Circle here
case Rectangle:
    fmt.Println(v.Width, v.Height)
default:
    fmt.Println("unknown shape")
}
```

## Key takeaways

- Interfaces are implicit — the compiler connects types and interfaces at the point of use
- `error` is just an interface — any type with `Error() string` satisfies it
- Use `errors.Is` for sentinel errors, `errors.As` for custom error types with data
- `%w` wraps errors so the chain stays inspectable — `%v` doesn't
- Only append valid values to a slice — validate at the boundary, keep downstream code clean
- `continue` exits a loop iteration, `break` just exits the switch case