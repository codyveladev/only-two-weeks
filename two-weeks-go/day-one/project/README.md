# go-calculator

A simple CLI calculator built in Go as part of a 14-day Go learning plan.

## Usage

```bash
go run main.go <num> <operator> <num>
```

```bash
go run main.go 10 + 5   # 15
go run main.go 10 - 3   # 7
go run main.go 6 * 7    # 42
go run main.go 10 / 2   # 5
go run main.go 10 % 3   # 1
```

## Supported operators

| Operator | Description |
|----------|-------------|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `%` | Modulo |

## Known limitations

- Integer only — no float support
- Integer division truncates (e.g. `7 / 2 = 3`)