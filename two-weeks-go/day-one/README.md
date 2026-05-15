# Day 1 — Hello, Go

Building a CLI calculator to learn the basics.

## What I learned

- Variables, constants, and basic types (`int`, `string`, `bool`)
- Functions with multiple return values
- `switch` statements
- `os.Args` for CLI arguments
- `strconv.Atoi` to parse strings to ints
- `fmt.Errorf` and the `error` type
- Error wrapping with `%w`

## What I built

A CLI calculator that takes two numbers and an operator as arguments, returns the result, and handles errors gracefully.

```bash
go run main.go 10 + 5   # Result is: 15
go run main.go 10 / 0   # cannot divide by zero
go run main.go 10 + abc # input must be a number: ...
go run main.go 10 +     # Usage: go run main.go <num> <operator> <num>
```

## Key takeaway

Go's idiomatic error handling — returning `(value, error)` and checking `err != nil` — is different from try/catch but keeps error handling explicit and close to where it happens.

## Known limitations

- Integer only, no float support
- Integer division truncates (`7 / 2 = 3`)